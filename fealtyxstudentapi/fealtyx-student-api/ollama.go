package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateSummary(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	mu.Lock()
	student, exists := students[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	prompt := "Generate a short summary for the following student:\n" +
		"Name: " + student.Name + "\n" +
		"Age: " + strconv.Itoa(student.Age) + "\n" +
		"Email: " + student.Email

	body, _ := json.Marshal(OllamaRequest{
		Model:  "llama3",
		Prompt: prompt,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to call Ollama", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result OllamaResponse
	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, &result)

	json.NewEncoder(w).Encode(map[string]string{"summary": result.Response})
}
