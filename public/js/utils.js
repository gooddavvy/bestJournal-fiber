// get and set year
let yearElement = document.getElementById("yearSpan");
let currentYear = new Date().getFullYear();

yearElement.textContent = currentYear;

// set the initial PIN and registration
let thePin = "0000";
let registration = {};

localStorage.setItem("thePin", thePin);