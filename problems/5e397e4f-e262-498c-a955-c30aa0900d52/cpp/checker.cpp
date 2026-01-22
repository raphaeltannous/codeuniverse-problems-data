#include "backend_checker.cpp"
#include "driver.cpp"
#include "solution.cpp"

namespace MainChecker {
    BackendChecker::RunResult runChecker(const std::vector<Driver::Testcase>& testcases) {
        BackendChecker::RunResult runResult;
        Solution solution;

        for (const auto& testcase : testcases) {
            std::string input = testcase.input;
            int expected = testcase.expected;
            
            auto capture = BackendChecker::captureStdout([&]() {
                return solution.romanToInteger(input);
            });
            
            int got = capture.result;

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
