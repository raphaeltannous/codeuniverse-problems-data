package main

func runChecker(testcases []*Testcase) *RunResult {
	runResult := new(RunResult)

	for _, testcase := range testcases {
		// input
		// expected
		// got

		if expected != got {

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
