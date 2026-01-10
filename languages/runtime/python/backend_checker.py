from driver import load_testcases
from checker import run_checker
import json
import time
import sys

try:
    from main import Solution
except SyntaxError as e:
    print(f"SyntaxError: {e.msg} at line {e.lineno}.\n\n{e.text}", file=sys.stderr)
    if e.offset:
        print(" " * (e.offset - 1) + "^", file=sys.stderr)
    sys.exit(1)
except Exception as e:
    print(f"{e.__class__.__name__}: {e}", file=sys.stderr)
    sys.exit(1)


def run_tests():
    testcases = load_testcases("testcases.json")
    solution = Solution()

    start_time = time.time()

    results = {}

    results = run_checker(testcases, solution)

    end_time = time.time()

    total_ms = (end_time - start_time) * 1000
    results["executionTime"] = total_ms

    if results.get("exception", False):
        print(f"{results['exception']}", file=sys.stderr)

    with open("results.json", "w") as file:
        json.dump(results, file)


if __name__ == "__main__":
    run_tests()
