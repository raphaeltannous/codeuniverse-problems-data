require_relative 'backend_checker'
require_relative 'solution'

def run_checker(testcases)
  run_result = BackendChecker::RunResult.new
  
  testcases.each do |testcase|
    input = testcase[:input]
    expected = testcase[:expected]
    
    capture = BackendChecker.capture_stdout { roman_to_integer(input) }
    got = capture[:result]
    
    if expected != got
      failed_testcase = BackendChecker::FailedTestcase.new(
        testcase[:id],
        testcase[:input],
        testcase[:expected],
        got,
        capture[:stdout]
      )
      
      run_result.failedTestcases << failed_testcase
      run_result.isPassed = false
    end
  end
  
  run_result
end
