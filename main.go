package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"statsbuddy/model"
)

// type Meta struct {
// 	DataVersion string `json:"data_version"`
// 	Created     string `json:"created"`
// 	Revision    int    `json:"revision"`
// }

// type Info struct {
// 	BallsPerOver  int       `json:"balls_per_over"`
// 	Venue         string    `json:"venue"`
// 	Dates         []string  `json:"dates"`
// 	Event         *Event    `json:"event,omitempty"`
// 	Gender        string    `json:"gender"`
// 	Teams         []string  `json:"teams"`
// 	Outcome       *Outcome  `json:"outcome"`
// 	Toss          Toss      `json:"toss"`
// 	PlayerOfMatch []string  `json:"player_of_match,omitempty"`
// 	MatchType     string    `json:"match_type"`
// 	Overs         int       `json:"overs"`
// 	Season        Season    `json:"season"`
// // 	Registry      Registry  `json:"registry"`
// // 	Officials     Officials `json:"officials"`
// // }

// // type Registry struct {
// // 	People map[string]string `json:"people"`
// // }

// type Season string

// func (s *Season) UnmarshalJSON(data []byte) error {
// 	var temp interface{}
// 	if err := json.Unmarshal(data, &temp); err != nil {
// 		return err
// 	}
// 	switch v := temp.(type) {
// 	case string:
// 		*s = Season(v)
// 	case float64:
// 		*s = Season(strconv.Itoa(int(v)))
// 	default:
// 		return fmt.Errorf("unexpected type for Season: %T", v)
// 	}
// 	return nil
// }

// type Event struct {
// 	Name        string `json:"name"`
// // 	MatchNumber int    `json:"match_number,omitempty"`
// 	Group       Group  `json:"group,omitempty"`
// // 	Stage       string `json:"stage,omitempty"`
// // }

// type Group string

// func (g *Group) UnmarshalJSON(data []byte) error {
// 	var temp interface{}
// 	if err := json.Unmarshal(data, &temp); err != nil {
// 		return err
// 	}
// 	switch v := temp.(type) {
// 	case string:
// 		*g = Group(v)
// 	case float64:
// 		*g = Group(strconv.Itoa(int(v)))
// 	default:
// 		return fmt.Errorf("unexpected type for Group: %T", v)
// 	}
// 	return nil
// }

// type Outcome struct {
// 	Winner string `json:"winner,omitempty"`
// 	By     *By    `json:"by,omitempty"`
// 	Result string `json:"result,omitempty"`
// 	Method string `json:"method,omitempty"`
// }

// type By struct {
// 	Runs    int `json:"runs,omitempty"`
// 	Wickets int `json:"wickets,omitempty"`
// 	Innings int `json:"innings,omitempty"`
// }

// type Toss struct {
// 	Decision    string `json:"decision"`
// 	Winner      string `json:"winner"`
// 	Uncontested bool   `json:"uncontested,omitempty"`
// }

// type Officials struct {
// 	MatchReferees []string `json:"match_referees,omitempty"`
// 	TVUmpires     []string `json:"tv_umpires,omitempty"`
// 	Umpires       []string `json:"umpires,omitempty"`
// }

// type Inning struct {
// 	Team   string  `json:"team"`
// 	Overs  []Over  `json:"overs"`
// 	Target *Target `json:"target,omitempty"`
// }

// type Over struct {
// 	Over       int        `json:"over"`
// 	Deliveries []Delivery `json:"deliveries"`
// }

// type Delivery struct {
// 	Batter     string   `json:"batter"`
// 	Bowler     string   `json:"bowler"`
// 	NonStriker string   `json:"non_striker"`
// 	Runs       Runs     `json:"runs"`
// 	Extras     *Extras  `json:"extras,omitempty"`
// 	Wickets    []Wicket `json:"wickets,omitempty"`
// 	Review     *Review  `json:"review,omitempty"`
// }

// type Runs struct {
// 	Batter      int  `json:"batter"`
// 	Extras      int  `json:"extras"`
// 	Total       int  `json:"total"`
// 	NonBoundary bool `json:"non_boundary,omitempty"`
// }

// type Extras struct {
// 	Byes    int `json:"byes,omitempty"`
// 	LegByes int `json:"legbyes,omitempty"`
// 	NoBalls int `json:"noballs,omitempty"`
// 	Wides   int `json:"wides,omitempty"`
// 	Penalty int `json:"penalty,omitempty"`
// }

// type Wicket struct {
// 	Kind      string    `json:"kind"`
// 	PlayerOut string    `json:"player_out"`
// 	Fielders  []Fielder `json:"fielders,omitempty"`
// }

// type Fielder struct {
// 	Name       string `json:"name"`
// 	Substitute bool   `json:"substitute,omitempty"`
// }

// type Review struct {
// 	By          string `json:"by"`
// 	Umpire      string `json:"umpire,omitempty"`
// 	Batter      string `json:"batter"`
// 	Decision    string `json:"decision"`
// 	UmpiresCall bool   `json:"umpires_call,omitempty"`
// }

// type Target struct {
// 	Overs float32 `json:"overs"`
// 	Runs  int     `json:"runs"`
// }

// type CricketMatch struct {
// 	Meta    Meta     `json:"meta"`
// 	Info    Info     `json:"info"`
// 	Innings []Inning `json:"innings"`
// }

func main() {
	// Path to the zip file
	// zipFilePath := "all_json.zip"

	// // Open the zip file
	// zipReader, err := zip.OpenReader(zipFilePath)
	// if err != nil {
	// 	log.Fatalf("Failed to open zip file: %v", err)
	// }
	// defer zipReader.Close()

	// // Store all parsed matches
	// var matches []model.CricketMatch

	// // Iterate through the files in the zip archive
	// for _, file := range zipReader.File {
	// 	fmt.Printf("Reading file: %s\n", file.Name)

	// 	// Open the file inside the zip
	// 	fileReader, err := file.Open()
	// 	if err != nil {
	// 		log.Printf("Failed to open file %s: %v", file.Name, err)
	// 		continue
	// 	}
	// 	defer fileReader.Close()

	// 	// Read the file content
	// 	content, err := ioutil.ReadAll(fileReader)
	// 	if err != nil {
	// 		log.Printf("Failed to read file %s: %v", file.Name, err)
	// 		continue
	// 	}

	// Path to the JSON files directory
	jsonDir := "india_male_json"

	// Store all parsed matches
	var matches []model.CricketMatch

	// Read all files from the directory
	files, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// Iterate through each file in the directory
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			fmt.Printf("Reading file: %s\n", file.Name())

			// Read the JSON file
			content, err := os.Open(filepath.Join(jsonDir, file.Name()))
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
				continue
			}
			defer content.Close()

			// Create a new JSON decoder
			decoder := json.NewDecoder(content)
			decoder.DisallowUnknownFields()

			// Parse the JSON content
			var match model.CricketMatch
			if err := decoder.Decode(&match); err != nil {
				log.Printf("Failed to parse JSON in file %s: %v", file.Name(), err)
				continue
			}

			// Append the parsed match to the list
			// matches = append(matches, match)
		}
	}
	// Print all parsed matches (for demonstration)
	for i, match := range matches {
		fmt.Printf("Match %d: %+v\n", i+1, match)
	}
}
