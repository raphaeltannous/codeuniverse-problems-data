package main

import "log"

func runChecker(testcases []*Testcase) *RunResult {
	runResult := &RunResult{
		IsPassed: true,
	}

	for _, testcase := range testcases {
		input := testcase.Input
		expected := testcase.Expected
		gotValue, stdout, err := captureStdout(isValid, input)
		if err != nil {
			log.Fatal(err)
		}

		got := gotValue[0].Bool()

		if expected != got {

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
