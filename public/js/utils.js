// get and set year
let yearElement = document.getElementById("yearSpan");
let currentYear = new Date().getFullYear();

yearElement.textContent = currentYear;

// set the initial PIN
let thePin = "0000";