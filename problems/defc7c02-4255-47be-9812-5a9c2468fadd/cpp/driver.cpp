using json = nlohmann::json;
using namespace std;

namespace Driver {
    struct Testcase {
        int id;
        int input;
        string expected;
        bool isPublic;
    };

    vector<Testcase> loadTestcases(const string& filename) {
        ifstream file(filename);
        if (!file.is_open()) {
            cerr << "Failed to open testcases file." << endl;
            exit(1);
        }

        json j;
        file >> j;

        vector<Testcase> testcases;
        for (const auto& item : j) {
            Testcase tc;
            tc.id = item["id"];
            tc.input = item["input"];
            tc.expected = item["expected"];
            tc.isPublic = item["isPublic"];
            testcases.push_back(tc);
        }

        return testcases;
    }
}
