// Register a new user
function registerUser() {
  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;

  fetch("/api/v1/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name: name, email: email }),
  })
    .then((response) => response.json())
    .then((data) => alert("User Registered: " + data.name))
    .catch((error) => console.error("Error:", error));
}

// Fetch all users
function getUsers() {
  fetch("/api/v1/get")
    .then((response) => response.json())
    .then((users) => {
      let userList = document.getElementById("userList");
      userList.innerHTML = "";
      users.forEach((user) => {
        let li = document.createElement("li");
        li.textContent = `ID: ${user.id}, Name: ${user.name}, Email: ${user.email}`;
        userList.appendChild(li);
      });
    })
    .catch((error) => console.error("Error:", error));
}

// Update user
function updateUser() {
  let id = document.getElementById("updateId").value;
  let name = document.getElementById("updateName").value;
  let email = document.getElementById("updateEmail").value;

  fetch(`/api/v1/update?id=${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name: name, email: email }),
  })
    .then((response) => {
      if (response.ok) {
        alert("User Updated Successfully!");
      } else {
        alert("User Not Found!");
      }
    })
    .catch((error) => console.error("Error:", error));
}

// Delete user
function deleteUser() {
  let id = document.getElementById("deleteId").value;

  fetch(`/api/v1/delete?id=${id}`, {
    method: "DELETE",
  })
    .then((response) => {
      if (response.ok) {
        alert("User Deleted Successfully!");
      } else {
        alert("User Not Found!");
      }
    })
    .catch((error) => console.error("Error:", error));
}
