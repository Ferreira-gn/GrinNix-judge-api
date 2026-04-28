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

type JudgeRequest struct {
	Code         string              `json:"code"`
	FunctionName string              `json:"function_name"`
	TestCases    []executor.TestCase `json:"test_cases"`
}

func main() {
	http.HandleFunc("/run", runHandler)
	
	
	http.HandleFunc("/judge", func(w http.ResponseWriter, r *http.Request) {
		var req JudgeRequest
		json.NewDecoder(r.Body).Decode(&req)

		results, err := executor.RunJS(req.Code, req.FunctionName, req.TestCases)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Calcula resumo
		passed := 0
		for _, r := range results {
			if r.Passed {
				passed++
			}
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"results":    results,
			"total":      len(results),
			"passed":     passed,
			"all_passed": passed == len(results),
		})
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
