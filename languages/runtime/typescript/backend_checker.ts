import * as fs from 'fs';

export class FailedTestcase {
    id: number;
    input: any;
    expected: any;
    got: any;
    stdOut: string;

    constructor(id: number, input: any, expected: any, got: any, stdOut: string) {
        this.id = id;
        this.input = input;
        this.expected = expected;
        this.got = got;
        this.stdOut = stdOut;
    }
}

export class RunResult {
    isPassed: boolean = true;
    failedTestcases: FailedTestcase[] = [];
    memoryUsage: number = 0.0;
    executionTime: number = 0.0;
    stdOut: string = '';
    stdErr: string = '';
}

export function captureStdout<T>(fn: () => T): { result: T; stdout: string } {
    const originalLog = console.log;
    let output = '';
    
    console.log = (...args: any[]) => {
        output += args.join(' ') + '\n';
    };
    
    const result = fn();
    
    console.log = originalLog;
    
    return { result, stdout: output };
}

export function writeResults(results: RunResult, filename: string): void {
    try {
        fs.writeFileSync(filename, JSON.stringify(results, null, 2));
    } catch (error) {
        console.error('Failed to write results file.');
        process.exit(1);
    }
}

export function runTests<T>(testcases: T[], runChecker: (testcases: T[]) => RunResult): void {
    const startTime = Date.now();
    
    const results = runChecker(testcases);
    
    const endTime = Date.now();
    results.executionTime = endTime - startTime;
    
    writeResults(results, 'results.json');
}
