package main

import (
	"encoding/json"
	"log"
	"os"
)

type Testcase struct {
	Id       int  `json:"id"`
	Input    CustomInput  `json:"input"`
	Expected [][]int  `json:"expected"`
	IsPublic bool `json:"isPublic"`
}

type CustomInput struct {
	List []int `json:"list"`
	K    int   `json:"k"`
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
