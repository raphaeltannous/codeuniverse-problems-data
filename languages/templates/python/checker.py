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
        input = testcase["input"]
        expected = testcase["expected"]
        result = {}

        try:
            # SET FUNCTION
            # got = solution.FUNCTION(input)

            if got != expected:
                result["id"] = testcase["id"]
                result["input"] = input
                result["expected"] = expected
                result["got"] = got

                results["isPassed"] = False
                results["failedTestcases"].append(result)
        except Exception as e:
            results["exception"] = e
            results["isPassed"] = False
            break

    return results
