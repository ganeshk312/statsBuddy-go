package api

import (
	"encoding/json"
	"net/http"
	"statsbuddy/service"
)

type ChatbotHandler struct {
	service *service.ChatbotService
}

func NewChatbotHandler(service *service.ChatbotService) *ChatbotHandler {
	return &ChatbotHandler{
		service: service,
	}
}

type QueryRequest struct {
	Prompt string `json:"prompt"`
}

type QueryResponse struct {
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (h *ChatbotHandler) HandleQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req QueryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.ProcessQuery(r.Context(), req.Prompt)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(QueryResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(QueryResponse{Response: response})
}
