package main

import (
	"log"
	"slices"
)

func runChecker(testcases []*Testcase) *RunResult {
	runResult := &RunResult{
		IsPassed: true,
	}

	for _, testcase := range testcases {
		input := sliceToLinkedList(testcase.Input)
		expected := testcase.Expected
		gotValue, stdout, err := captureStdout(addOne, input)
		if err != nil {
			log.Fatal(err)
		}

		got := linkedListToSlice(gotValue[0].Interface().(*Node))

		if !slices.Equal(expected, got) {
			runResult.FailedTestcases = append(runResult.FailedTestcases, &FailedTestcase{
				Id:       testcase.Id,
				Input:    testcase.Input,
				Expected: testcase.Expected,
				Got:      got,
				StdOut:   stdout,
			})
			runResult.IsPassed = false
		}
	}

	return runResult
}
