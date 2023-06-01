let btn = document.querySelector("button");
localStorage.setItem("registered", "false");
let registered = localStorage.getItem("registered");

// Check if registered

if (registered) {
    location.href = "/home.html";
}

// Open popup

btn.onclick = () => {
    window.open("/register.html");
    localStorage.setItem("registered", true)
    registered = localStorage.getItem("registered");
    location.reload();
}