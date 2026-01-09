from driver import load_testcases
from checker import run_checker
import time
import sys

try:
    from main import Solution
except SyntaxError as e:
    print(f"SyntaxError: {e.msg} at line {e.lineno}.\n\n{e.text}", end="")
    if e.offset:
        print(" " * (e.offset - 1) + "^")
    sys.exit(1)
except Exception as e:
    print(f"{e.__class__.__name__}: {e}")
    sys.exit(1)


def run_tests():
    testcases = load_testcases()
    solution = Solution()

    start_time = time.time()

    results = {}

    results = run_checker(testcases, solution)

    end_time = time.time()

    total_ms = (end_time - start_time) * 1000
    print(total_ms)

    # TODO write results as a json file
    return results


if __name__ == "__main__":
    run_tests()
