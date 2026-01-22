#include "backend_checker.cpp"
#include "driver.cpp"
#include "solution.cpp"

namespace MainChecker {
    BackendChecker::RunResult runChecker(const std::vector<Driver::Testcase>& testcases) {
        BackendChecker::RunResult runResult;
        Solution solution;

        for (const auto& testcase : testcases) {
            auto input = testcase.input;
            auto expected = testcase.expected;
            
            auto capture = BackendChecker::captureStdout([&]() {
                return solution.funcName(input);
            });
            
            auto got = capture.result;

            if (expected != got) {
                BackendChecker::FailedTestcase failedTestcase;
                failedTestcase.id = testcase.id;
                failedTestcase.input = testcase.input;
                failedTestcase.expected = testcase.expected;
                failedTestcase.got = got;
                failedTestcase.stdOut = capture.stdout;

                runResult.failedTestcases.push_back(failedTestcase);
                runResult.isPassed = false;
            }
        }

        return runResult;
    }
}
