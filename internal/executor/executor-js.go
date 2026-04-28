package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)
type TestCase struct {
    Input    interface{} `json:"input"`
    Expected interface{} `json:"expected"`
}

type TestResult struct {
    Index    int         `json:"index"`
    Passed   bool        `json:"passed"`
    Output   interface{} `json:"output,omitempty"`
    Expected interface{} `json:"expected,omitempty"`
    Error    string      `json:"error,omitempty"`
}

func RunJS(code, functionName string, testCases []TestCase) ([]TestResult, error) {
    tcJSON, _ := json.Marshal(testCases)

    cmd := exec.Command("docker", "run", "--rm",
        "--network", "none",
        "--memory", "128m",
        "--cpus", "0.5",
        "-e", "USER_CODE="+code,
        "-e", "FUNCTION_NAME="+functionName,
        "-e", "TEST_CASES="+string(tcJSON),
        "-e", "TIMEOUT=5",
        "js-runner",
    )

    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Run(); err != nil {
        // timeout ou crash — retorna erro geral
        return nil, fmt.Errorf("execution failed: %s", stderr.String())
    }

    var results []TestResult
    if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
        return nil, fmt.Errorf("bad output: %w", err)
    }

    return results, nil
}