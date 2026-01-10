package main

func runChecker(testcases []*Testcase) *RunResult {
	runResult := &RunResult{
		IsPassed: true,
	}

	for _, testcase := range testcases {
		input := testcase.Input
		expected := testcase.Expected
		got := climbStairs(input)

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
