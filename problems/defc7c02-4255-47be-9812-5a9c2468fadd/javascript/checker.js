const { FailedTestcase, RunResult, captureStdout } = require('./backend_checker.js');
const { integerToRoman } = require('./solution.js');

function runChecker(testcases) {
    const runResult = new RunResult();
    
    for (const testcase of testcases) {
        const input = testcase.input;
        const expected = testcase.expected;
        
        const capture = captureStdout(() => integerToRoman(input));
        const got = capture.result;
        
        if (expected !== got) {
            const failedTestcase = new FailedTestcase(
                testcase.id,
                testcase.input,
                testcase.expected,
                got,
                capture.stdout
            );
            
            runResult.failedTestcases.push(failedTestcase);
            runResult.isPassed = false;
        }
    }
    
    return runResult;
}

module.exports = { runChecker };
