const { runTests } = require('./backend_checker.js');
const { loadTestcases } = require('./driver.js');
const { runChecker } = require('./checker.js');

const testcases = loadTestcases('testcases.json');
runTests(testcases, runChecker);
