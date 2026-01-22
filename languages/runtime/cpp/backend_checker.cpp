#ifndef BACKEND_CHECKER_CPP
#define BACKEND_CHECKER_CPP

#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <chrono>
#include <functional>
#include "json.hpp"

using json = nlohmann::json;

namespace BackendChecker {
    struct FailedTestcase {
        int id;
        json input;
        json expected;
        json got;
        std::string stdOut;

        json toJson() const {
            return {
                {"id", id},
                {"input", input},
                {"expected", expected},
                {"got", got},
                {"stdOut", stdOut}
            };
        }
    };

    struct RunResult {
        bool isPassed = true;
        std::vector<FailedTestcase> failedTestcases;
        double memoryUsage = 0.0;
        double executionTime = 0.0;
        std::string stdOut;
        std::string stdErr;

        json toJson() const {
            json failedArray = json::array();
            for (const auto& tc : failedTestcases) {
                failedArray.push_back(tc.toJson());
            }

            return {
                {"isPassed", isPassed},
                {"failedTestcases", failedArray},
                {"memoryUsage", memoryUsage},
                {"executionTime", executionTime},
                {"stdOut", stdOut},
                {"stdErr", stdErr}
            };
        }
    };

    template<typename T>
    struct StdoutCapture {
        T result;
        std::string stdout;
    };

    template<typename Func>
    auto captureStdout(Func fn) -> StdoutCapture<decltype(fn())> {
        std::stringstream buffer;
        std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
        
        auto result = fn();
        
        std::cout.rdbuf(old);
        
        return {result, buffer.str()};
    }

    void writeResults(const RunResult& results, const std::string& filename) {
        std::ofstream file(filename);
        if (!file.is_open()) {
            std::cerr << "Failed to write results file." << std::endl;
            exit(1);
        }
        file << results.toJson().dump();
        file.close();
    }

    template<typename T, typename Func>
    void runTests(const std::vector<T>& testcases, Func runChecker) {
        auto startTime = std::chrono::high_resolution_clock::now();

        RunResult results = runChecker(testcases);

        auto endTime = std::chrono::high_resolution_clock::now();
        auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(endTime - startTime);
        results.executionTime = static_cast<double>(duration.count());

        writeResults(results, "results.json");
    }
}

#endif
