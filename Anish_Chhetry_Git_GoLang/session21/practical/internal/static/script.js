function createUser () {
    const name = document.getElementById('createName').value;
    fetch('/api/v1/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: name }),
    })
        .then(response => response.json())
        .then(data => {
            alert(data.message + "\n ID: " + data.id +"\n Name: "+name);
        })
        .catch(error => console.error('Error:', error));
}

function getUsers() {
    fetch('/api/v1/get')
        .then(response => response.json())
        .then(data => {
            const tableBody = document.querySelector('#userTable tbody');
            tableBody.innerHTML = '';
            if (data && data.users) {
                data.users.forEach(user => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${user.id !== undefined ? user.id : 'N/A'}</td>
                        <td>${user.name !== undefined ? user.name : 'N/A'}</td>
                    `;
                    tableBody.appendChild(row);
                });
            }
        })
        .catch(error => console.error('Error:', error));
}


function getUserById(id) {
    fetch(`/api/v1/get/${id}`)
        .then(response => response.json())
        .then(data => {
            if (data.user) {
                alert(`User ID: ${data.user.id}\nName: ${data.user.name}`);
            } else {
                alert('User not found.');
            }
        })
        .catch(error => console.error('Error:', error));
}

function updateUser () {
    const id = document.getElementById('updateId').value;
    const name = document.getElementById('updateName').value;
    fetch(`/api/v1/update/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: name }),
    })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
        })
        .catch(error => console.error('Error:', error));
}

function deleteUser () {
    const id = document.getElementById('deleteId').value;
    fetch(`/api/v1/delete/${id}`, {
        method: 'DELETE',
    })
        .then(response => response.json())
        .then(data => {
            alert(data.message);
        })
        .catch(error => console.error('Error:', error));
}
