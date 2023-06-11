let app = document.getElementById("App");
let form = document.querySelector("form");

let err = null;
let isSuccess = false; // Flag to indicate success

// form data
let formData = {};

// functions
function updateFormData(data, title, shortDesc) {
    if (title !== "" && shortDesc !== "") {
        data.title = title;
        data.shortDesc = shortDesc;
    } else {
        throw new Error("RequestError: Each input provided in the form must be filled in, not left empty.");
    }
}

// form submit
form.addEventListener("submit", e => {
    e.preventDefault();

    let form_title = document.getElementById("pageTitle").value;
    let form_shortDesc = document.getElementById("pageShortDesc").value;

    try {
        updateFormData(formData, form_title, form_shortDesc);
    } catch (error) {
        err = error;
    }

    form.reset();
    app.innerHTML = `
        <h1>Please wait</h1>
        <p>Processing provided form data...</h1>
        <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">Loading...</span>
        </div>
    `;

    window.setTimeout(() => {
        if (err instanceof Error) {
            let errMessage = err.message;
            app.innerHTML = `
                <h1>Error</h1>
                <p>An error occurred creating the journal page:</p>
                <p class="text-danger"><code>${errMessage}</code></p>
                <button onclick="location.reload()" class="btn btn-primary">Retry</button>
            `;
            console.error(errMessage);
        } else {
            (async () => {
                try {
                    let response = await fetch(`https://${location.host}/api/addPage?title=${formData.title}&desc=${formData.shortDesc}`, {
                        method: "GET",
                        cache: "no-store",
                    });

                    if (!response.ok) {
                        throw new Error("Error: " + response.status);
                    }

                    const data = await response.json();
                    console.log(data);
                    if (data.message === "Successfully added journal page!") {
                        alert("Added journal page. We're good to go.");
                        isSuccess = true; // Set the flag to indicate success
                    }
                } catch (error) {
                    err = error;
                    let errMessage = err.message;
                    app.innerHTML = `
                        <h1>Error</h1>
                        <p>An error occurred looking in the journal page server:</p>
                        <p class="text-danger"><code>${errMessage}</code></p>
                        <button onclick="location.reload()" class="btn btn-primary">Retry</button>
                    `;
                    throw err;
                }
            })();
        }
    }, 2500);
});

// custom submit function
function submitForm() {
    let submitEvent = new Event("submit", {
        bubbles: true,
        cancelable: true
    });
    form.dispatchEvent(submitEvent);
}

// button click event
let submitButton = document.querySelector("form button");
submitButton.addEventListener("click", e => {
    e.preventDefault();
    submitForm();
});

// check for success

if (isSuccess) {
    app.innerHTML = `
        <h1>Success creating page</h1>
        <p>You'll be redirected shortly.</p>
    `;
    window.setTimeout(() => (location.href = "/home.html"), 800);
}