const fs = require('fs');

function loadTestcases(filename) {
    try {
        const data = fs.readFileSync(filename, 'utf8');
        return JSON.parse(data);
    } catch (error) {
        console.error('Failed to open testcases file.');
        process.exit(1);
    }
}

module.exports = { loadTestcases };
