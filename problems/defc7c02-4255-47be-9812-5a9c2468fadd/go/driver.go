package main

import (
	"encoding/json"
	"log"
	"os"
)

// Update:
// Input any accordingly
// Expected any accordingly
type Testcase struct {
	Id       int  `json:"id"`
	Input    any  `json:"input"`
	Expected any  `json:"expected"`
	IsPublic bool `json:"isPublic"`
}

func loadTestcases(filename string) []*Testcase {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open testcases file.")
	}

	var testcases []*Testcase

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&testcases)
	if err != nil {
		log.Fatal("failed to decode testcases file.")
	}

	return testcases
}
