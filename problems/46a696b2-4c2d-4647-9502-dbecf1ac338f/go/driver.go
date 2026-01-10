package main

import (
	"encoding/json"
	"log"
	"os"
)

type Testcase struct {
	Id       int  `json:"id"`
	Input    []int  `json:"input"`
	Expected []int  `json:"expected"`
	IsPublic bool `json:"isPublic"`
}

type Node struct {
	Val int
	Next *Node
}

func sliceToLinkedList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	dummy := &Node{}
	curr := dummy

	for _, num := range nums {
		curr.Next = &Node{Val: num}
		curr = curr.Next
	}

	return dummy.Next
}

func linkedListToSlice(head *Node) []int {
	nums := make([]int, 0, 50)
	for e := head; e != nil; e = e.Next {
		nums = append(nums, e.Val)
	}
	return nums
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
