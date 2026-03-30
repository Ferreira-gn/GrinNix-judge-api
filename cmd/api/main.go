package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ferreira-gn/judge-api/internal/executor"
)

type Request struct {
	Code string `json:"code"`
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	result := executor.RunTypeScript(req.Code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/run", runHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
