let app = document.getElementById("App");

// Fetch and render the journal pages
// Fetch and render the journal pages
(async () => {
    try {
        // Fetch the journal pages
        let response = await fetch(`/api/journalPages`, {
            method: "GET",
            cache: "no-store",
        });

        if (!response.ok) {
            throw new Error("Error: " + response.status);
        }

        const data = await response.json();
        journalPages = data;

        console.log(journalPages);

        // Render the journal pages
        let cardContent = document.querySelector(".cardContent");
        cardContent.innerHTML = ""; // Clear existing content

        if (journalPages.length > 0) {
            journalPages.forEach(journalPage => {
                const { title, date, shortDesc, link } = journalPage;
                if (title && date && shortDesc && link) {
                    const cardHTML = `
                        <div class="card" style="width: 100%;">
                            <div class="card-body">
                                <h5 class="card-title">${title}</h5>
                                <h6 class="card-subtitle mb-2 text-body-secondary">Date: <b>${date}</b></h6>
                                <p class="card-text">${shortDesc}</p>
                                <a href="${link}" class="btn btn-primary">Open</a>
                            </div>
                        </div>
                    `;
                    cardContent.innerHTML += cardHTML;
                }
            });
        } else {
            app.innerHTML = `
                <h1>All caught up!</h1>
                <p>You don't have any journal pages yet.</p>
                <button type="button" class="btn btn-primary">Go make one!</button>
            `;
        }
    } catch (err) {
        app.innerHTML = `
            <p>An error occurred:</p>
            <p style="color: red"><code>${err}</code></p>
            <button type="button" class="btn btn-primary" onclick="location.reload()">Reload</button>
        `;
        throw err;
    }
})();
