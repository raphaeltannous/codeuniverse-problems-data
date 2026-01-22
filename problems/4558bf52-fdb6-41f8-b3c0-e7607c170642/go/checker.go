package main

import (
	"log"
	"reflect"
)

func runChecker(testcases []*Testcase) *RunResult {
	runResult := &RunResult{
		IsPassed: true,
	}

	for _, testcase := range testcases {
		input := testcase.Input
		expected := testcase.Expected
		gotValue, stdout, err := captureStdout(splitLinkedList, sliceToLinkedList(input.List), input.K)
		if err != nil {
			log.Fatal(err)
		}

		got := sliceOfNodesToLinkedList(gotValue[0].Interface().([]*Node))

		if !reflect.DeepEqual(expected, got) {

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
