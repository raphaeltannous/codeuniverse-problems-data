package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type RunResult struct {
	IsPassed bool `json:"isPassed"`

	FailedTestcases []*FailedTestcase `json:"failedTestcases"`

	MemoryUsage   float64 `json:"memoryUsage"`
	ExecutionTime float64 `json:"executionTime"`

	StdOut string `json:"stdOut"`
	StdErr string `json:"stdErr"`
}

type FailedTestcase struct {
	Id       int `json:"id"`
	Input    any `json:"input"`
	Expected any `json:"expected"`
	Got      any `json:"got"`
}

func main() {
	testcases := loadTestcases("testcases.json")

	startTime := time.Now()

	var results *RunResult
	results = runChecker(testcases)

	results.ExecutionTime = float64(time.Since(startTime).Milliseconds())

	resultsFile, err := os.Create("results.json")
	if err != nil {
		log.Fatal("failed to create resultsFile.")
	}

	encoder := json.NewEncoder(resultsFile)
	if err := encoder.Encode(results); err != nil {
		log.Fatal("failed to write resultsFile.")
	}
}
