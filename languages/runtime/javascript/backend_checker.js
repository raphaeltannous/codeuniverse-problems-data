const fs = require('fs');

class FailedTestcase {
    constructor(id, input, expected, got, stdOut) {
        this.id = id;
        this.input = input;
        this.expected = expected;
        this.got = got;
        this.stdOut = stdOut;
    }
}

class RunResult {
    constructor() {
        this.isPassed = true;
        this.failedTestcases = [];
        this.memoryUsage = 0.0;
        this.executionTime = 0.0;
        this.stdOut = '';
        this.stdErr = '';
    }
}

function captureStdout(fn) {
    const originalLog = console.log;
    let output = '';
    
    console.log = (...args) => {
        output += args.join(' ') + '\n';
    };
    
    const result = fn();
    
    console.log = originalLog;
    
    return { result, stdout: output };
}

function writeResults(results, filename) {
    try {
        fs.writeFileSync(filename, JSON.stringify(results, null, 2));
    } catch (error) {
        console.error('Failed to write results file.');
        process.exit(1);
    }
}

function runTests(testcases, runChecker) {
    const startTime = Date.now();
    
    const results = runChecker(testcases);
    
    const endTime = Date.now();
    results.executionTime = endTime - startTime;
    
    writeResults(results, 'results.json');
}

module.exports = { FailedTestcase, RunResult, captureStdout, writeResults, runTests };
