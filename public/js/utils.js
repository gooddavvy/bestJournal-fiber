console.log(`    Interested in contributing?
    Go to https://github.com/gooddavvy/bestJournal-fiber/blob/main/CONTRIBUTE.md.
    THIS IS NOT A SCAM!!1
`);

// get and set year
let yearElement = document.getElementById("yearSpan");
let currentYear = new Date().getFullYear();

try { yearElement.textContent = currentYear; } catch { }

// set the initial PIN and registration
let thePin = "0000";
let registration = {};

localStorage.setItem("thePin", thePin);