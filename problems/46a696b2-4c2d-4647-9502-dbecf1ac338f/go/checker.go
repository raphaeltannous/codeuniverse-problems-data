package main

import "slices"

func runChecker(testcases []*Testcase) *RunResult {
	runResult := &RunResult{
		IsPassed: true,
	}

	for _, testcase := range testcases {
		input := sliceToLinkedList(testcase.Input)
		expected := testcase.Expected
		got := linkedListToSlice(addOne(input))

		if !slices.Equal(expected, got) {
			runResult.FailedTestcases = append(runResult.FailedTestcases, &FailedTestcase{
				Id:       testcase.Id,
				Input:    testcase.Input,
				Expected: testcase.Expected,
				Got:      got,
			})
			runResult.IsPassed = false
		}
	}

	return runResult
}
