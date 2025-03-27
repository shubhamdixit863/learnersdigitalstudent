const API_URL = "http://localhost:8080";

// Function to register a user
function registerUser() {
    const name = prompt("Enter Name:");
    const age = prompt("Enter Age:");

    if (!name || !age) return alert("Name and Age are required!");

    fetch(`${API_URL}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, age: parseInt(age) })
    })
        .then(res => res.json())
        .then(data => {
            alert(`User Registered: ID ${data.id}, Name: ${data.name}`);
            getUsers(); // Refresh user list after registration
        })
        .catch(err => console.error("Error:", err));
}

// Function to update a user
function updateUser(id) {
    if (!id) id = prompt("Enter User ID to Update:");
    const name = prompt("Enter New Name:");
    const age = prompt("Enter New Age:");

    if (!id || !name || !age) return alert("ID, Name, and Age are required!");

    fetch(`${API_URL}/update?id=${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, age: parseInt(age) })
    })
        .then(res => res.json())
        .then(data => {
            alert(`User Updated: ID ${data.id}, Name: ${data.name}`);
            getUsers(); // Refresh user list after update
        })
        .catch(err => console.error("Error:", err));
}

// Function to get all users and display them in a table
function getUsers() {
    fetch(`${API_URL}/get`)
        .then(res => res.json())
        .then(users => {
            const tableBody = document.querySelector("#userList tbody");
            tableBody.innerHTML = ""; // Clear table before adding new users

            for (const id in users) {
                const user = users[id];

                let row = document.createElement("tr");
                row.innerHTML = `
                <td>${user.id}</td>
                <td>${user.name}</td>
                <td>${user.age}</td>
                <td class="action-buttons">
                    <button onclick="updateUser(${user.id})">
                        <i class="fas fa-edit"></i> Edit
                    </button>
                    <button onclick="deleteUser(${user.id})">
                        <i class="fas fa-trash"></i> Delete
                    </button>
                </td>
            `;
                tableBody.appendChild(row);
            }
        })
        .catch(err => console.error("Error:", err));
}

// Function to delete a user
function deleteUser(id) {
    if (!id) id = prompt("Enter User ID to Delete:");

    if (!id) return alert("User ID is required!");

    fetch(`${API_URL}/delete?id=${id}`, { method: "DELETE" })
        .then(() => {
            alert("User Deleted");
            getUsers(); // Refresh user list after deletion
        })
        .catch(err => console.error("Error:", err));
}

// Load users when page loads
window.onload = getAllUsers;
