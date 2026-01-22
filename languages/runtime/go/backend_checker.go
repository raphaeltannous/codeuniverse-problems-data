package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	_ "time/tzdata"
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
	StdOut   any `json:"stdOut"`
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

func captureStdout(fn any, args ...any) ([]reflect.Value, string, error) {
	funcValue := reflect.ValueOf(fn)

	if funcValue.Kind() != reflect.Func {
		return nil, "", errors.New("Not a function")
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	arguments := make([]reflect.Value, len(args))
	for i, arg := range args {
		arguments[i] = reflect.ValueOf(arg)
	}

	returns := funcValue.Call(arguments)

	w.Close()
	os.Stdout = old

	var buf strings.Builder
	io.Copy(&buf, r)

	return returns, buf.String(), nil
}
