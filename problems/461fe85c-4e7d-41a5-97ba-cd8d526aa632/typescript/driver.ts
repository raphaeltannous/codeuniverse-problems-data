import * as fs from 'fs';

interface Testcase {
    id: number;
    input: string;
    expected: boolean;
    isPublic: boolean;
}

export function loadTestcases(filename: string): Testcase[] {
    try {
        const data = fs.readFileSync(filename, 'utf8');
        return JSON.parse(data);
    } catch (error) {
        console.error('Failed to open testcases file.');
        process.exit(1);
    }
}

export type { Testcase };
