#include "backend_checker.cpp"
#include "checker.cpp"

int main() {
    std::vector<Driver::Testcase> testcases = Driver::loadTestcases("testcases.json");
    BackendChecker::runTests(testcases, MainChecker::runChecker);
    return 0;
}
