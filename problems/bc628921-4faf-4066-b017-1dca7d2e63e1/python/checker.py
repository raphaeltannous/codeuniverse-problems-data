import io
from contextlib import redirect_stdout

def run_checker(testcases, solution):
    # results:
    # - exception
    # - isPassed
    # - failedTestcases:
    #      - id
    #      - input
    #      - expected
    #      - got
    #      - stdOut
    results = {
        "isPassed": True,
        "failedTestcases": [],
    }

    for testcase in testcases:
        testcase_input = testcase["input"]
        expected = testcase["expected"]
        result = {}

        try:
            stdout = io.StringIO()
            with redirect_stdout(stdout):
                got = solution.toLower(testcase_input)

            if got != expected:
                result["id"] = testcase["id"]
                result["input"] = testcase_input
                result["expected"] = expected
                result["got"] = got
                result["stdout"] = stdout.getvalue()

                results["isPassed"] = False
                results["failedTestcases"].append(result)
        except Exception as e:
            results["exception"] = f"{e}"
            results["isPassed"] = False
            break

    return results
