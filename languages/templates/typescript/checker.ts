import { FailedTestcase, RunResult, captureStdout } from './backend_checker';
import { Testcase } from './driver';
import { funcName } from './solution';

export function runChecker(testcases: Testcase[]): RunResult {
    const runResult = new RunResult();
    
    for (const testcase of testcases) {
        const input = testcase.input;
        const expected = testcase.expected;
        
        const capture = captureStdout(() => funcName(input));
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
