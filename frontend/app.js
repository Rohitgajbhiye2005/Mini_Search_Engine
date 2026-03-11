async function search() {
    const query = document.getElementById("searchQuery").value;
    if (!query) return;

    try {
        const res = await fetch(`http://localhost:8080/search?q=${encodeURIComponent(query)}&limit=10`);
        const data = await res.json();

        const resultsDiv = document.getElementById("results");
        resultsDiv.innerHTML = "";

        data.forEach(item => {
            const div = document.createElement("div");
            div.className = "result-item";
            div.innerHTML = `<a href="${item.URL}" target="_blank">${item.Title}</a><p>${item.URL}</p>`;
            resultsDiv.appendChild(div);
        });
    } catch (err) {
        console.error("Error fetching search results:", err);
    }
}

// attach search function to form
document.getElementById("searchForm").addEventListener("submit", (e) => {
    e.preventDefault();
    search();
});