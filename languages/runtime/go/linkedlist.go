package main

type Node struct {
	Val  int
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

func sliceOfNodesToLinkedList(lists []*Node) [][]int {
	result := make([][]int, len(lists))

	for i, numbers := range lists {
		result[i] = linkedListToSlice(numbers)
	}

	return result
}
