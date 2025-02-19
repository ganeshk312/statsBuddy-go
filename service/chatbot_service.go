package service

import (
	"context"
	"fmt"
	"log"
	"statsbuddy/config"
	"statsbuddy/llm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatbotService struct {
	llmProvider llm.LLMProvider
	mongoClient *mongo.Client
	config      *config.Config
}

func NewChatbotService(config *config.Config, mongoClient *mongo.Client) (*ChatbotService, error) {
	factory := &llm.LLMFactory{}
	provider, err := factory.CreateLLM(llm.LLMType(config.LLMType))
	if err != nil {
		return nil, fmt.Errorf("failed to create LLM provider: %v", err)
	}

	return &ChatbotService{
		llmProvider: provider,
		mongoClient: mongoClient,
		config:      config,
	}, nil
}

func (s *ChatbotService) ProcessQuery(ctx context.Context, userPrompt string) (string, error) {
	// Step 1: Generate MongoDB query using LLM
	mongoQuery, err := s.llmProvider.GenerateMongoQuery(userPrompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate MongoDB query: %v", err)
	}
	log.Println("Generated MongoDB query:", mongoQuery)

	// Step 2: Execute MongoDB query
	var result map[string]interface{}
	collection := s.mongoClient.Database(s.config.DatabaseName).Collection("matches")

	// Parse the MongoDB query string into BSON
	// var pipeline []bson.M
	// if err := bson.UnmarshalExtJSON([]byte(mongoQuery), true, &pipeline); err != nil {
	// 	return "", fmt.Errorf("failed to parse MongoDB query: %v", err)
	// }
	// pipeline := []bson.M{
	// 	{"$match": bson.M{"info.player_of_match": "JJ Bumrah"}},
	// 	{"$sort": bson.M{"info.dates": -1}},
	// 	{"$limit": 5},
	// 	{"$project": bson.M{"_id": 0, "teams": "$info.teams"}},
	// 	{"$unwind": "$teams"},
	// 	{"$match": bson.M{"teams": bson.M{"$ne": "JJ Bumrah"}}},
	// 	{"$group": bson.M{
	// 		"_id":       nil,
	// 		"opponents": bson.M{"$addToSet": "$teams"},
	// 	}},
	// 	{"$project": bson.M{"_id": 0, "opponents": 1}},
	// }

	// pipelineJSON := `[
	// 	{ "$match": { "info.player_of_match": "JJ Bumrah" } },
	// 	{ "$sort": { "info.dates": -1 } },
	// 	{ "$limit": 5 },
	// 	{ "$project": {
	// 	  "_id": 0,
	// 	  "opponent": { "$arrayElemAt": [ { "$setDifference": [ "$info.teams", ["India"] ] }, 0 ] }
	// 	}}
	//   ]`

	// Convert JSON string to BSON
	var pipeline mongo.Pipeline
	err = bson.UnmarshalExtJSON([]byte(mongoQuery), true, &pipeline)
	if err != nil {
		log.Fatal(err)
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return "", fmt.Errorf("failed to execute MongoDB query: %v", err)
	}
	defer cursor.Close(ctx)

	// Collect all results
	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		return "", fmt.Errorf("failed to decode MongoDB results: %v", err)
	}

	if len(results) == 0 {
		return "No results found for your query.", nil
	}

	result = map[string]interface{}{
		"count": len(results),
		"data":  results,
	}

	// Step 3: Format the response using LLM
	response, err := s.llmProvider.FormatResponse(result, userPrompt)
	if err != nil {
		return "", fmt.Errorf("failed to format response: %v", err)
	}

	return response, nil
}
