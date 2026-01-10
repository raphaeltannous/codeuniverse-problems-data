def run_checker(testcases, solution):
    # results:
    # - exception
    # - isPassed
    # - failedTestcases:
    #      - id
    #      - input
    #      - expected
    #      - got
    results = {
        "isPassed": True,
        "failedTestcases": [],
    }

    for testcase in testcases:
        testcase_input = testcase["input"]
        expected = testcase["expected"]
        result = {}

        try:
            got = solution.climbStairs(testcase_input)

            if got != expected:
                result["id"] = testcase["id"]
                result["input"] = testcase_input
                result["expected"] = expected
                result["got"] = got

                results["isPassed"] = False
                results["failedTestcases"].append(result)
        except Exception as e:
            results["exception"] = e
            results["isPassed"] = False
            break

    return results
