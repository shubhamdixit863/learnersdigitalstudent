document.getElementById("registerForm").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the default form submission

    const name = document.getElementById("name").value;
    const userId = document.getElementById("userId").value; // Get the user ID

    fetch("/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ name: name, id: parseInt(userId) }), // Send both name and ID
    })
        .then(response => response.text())
        .then(data => {
            document.getElementById("response").innerText = data;
        })
        .catch(error => console.error("Error:", error));
});

function updateUser () {
    const userId = document.getElementById("userId").value; // Get the user ID
    const name = document.getElementById("name").value; // Get the updated name
    if (!userId) {
        alert("Please enter a User ID");
        return;
    }

    fetch(`/update?id=${userId}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ name: name }), // Send the updated name
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok " + response.statusText);
            }
            return response.text();
        })
        .then(data => {
            document.getElementById("response").innerText = data;
        })
        .catch(error => console.error("Error:", error));
}

function getUser () {
    const userId = document.getElementById("userId").value;
    let url = "/get"; // Base URL

    if (!userId) {
        alert("Please enter a User ID");
        return;
    }
    url += `?id=${userId}`; // Append user ID if provided


    fetch(url)
        .then(response => response.text())
        .then(data => {
            document.getElementById("response").innerText = data;
        })
        .catch(error => console.error("Error:", error));
}

function deleteUser () {
    const userId = document.getElementById("userId").value;
    if (!userId) {
        alert("Please enter a User ID");
        return;
    }

    fetch(`/delete?id=${userId}`, { method: "DELETE" })
        .then(response => response.text())
        .then(data => {
            document.getElementById("response").innerText = data;
        })
        .catch(error => console.error("Error:", error));
}

function getAllUser(){
    let url = "/get"; // Base URL

    fetch(url)
        .then(response => response.text())
        .then(data => {
            document.getElementById("response").innerText = data;
        })
        .catch(error => console.error("Error:", error));
}