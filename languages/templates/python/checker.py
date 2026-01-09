def run_checker(testcases, solution):
    # result:
    #  - id
    #  - input
    #  - expected
    #  - got
    #  - exception
    #  - isPassed
    result = {
        "isPassed": True,
    }

    for testcase in testcases:
        input = testcase["input"]
        expected = testcase["expected"]

        try:
            # SET FUNCTION
            # got = solution.FUNCTION(input)

            if got != expected:
                result.id = testcase["id"]
                result.input = input
                result.expected = expected
                result.got = got
                result.isPassed = False
                break
        except Exception as e:
            result.exception = e
            result.isPassed = False
            break

    return result
