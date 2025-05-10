async function getUsers() {
    let response = await fetch('/get');
    let users = await response.json();
    let usersList = document.getElementById('users-list');
    usersList.innerHTML = '<h2>Users:</h2>' + users.map(user => `<p>${user.id}: ${user.name} (${user.email})</p>`).join('');
}

function showForm(type) {
    let form = `<input id="name" placeholder="Name"><input id="email" placeholder="Email"><button onclick="registerUser()">Submit</button>`;
    document.getElementById('form-container').innerHTML = form;
}

async function registerUser() {
    let name = document.getElementById('name').value;
    let email = document.getElementById('email').value;
    await fetch('/register', { method: 'POST', body: JSON.stringify({ name, email }) });
    getUsers();
}
