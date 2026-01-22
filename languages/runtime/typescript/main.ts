import { runTests } from './backend_checker';
import { loadTestcases } from './driver';
import { runChecker } from './checker';

const testcases = loadTestcases('testcases.json');
runTests(testcases, runChecker);
