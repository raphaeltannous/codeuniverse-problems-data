require_relative 'backend_checker'
require_relative 'driver'
require_relative 'checker'

testcases = load_testcases('testcases.json')
BackendChecker.run_tests(testcases, method(:run_checker))
