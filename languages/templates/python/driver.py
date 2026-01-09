import json


def load_tests(testcases_file: "testcases.json"):
    with open(testcases_file, "r", encoding="utf-8") as file:
        return json.load(file)
