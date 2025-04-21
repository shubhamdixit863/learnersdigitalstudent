import { useState, useEffect } from "react";
import axios from "../api/axios";

const AdminDashboard = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    axios.get("/admin/users").then((res) => setUsers(res.data));
  }, []);

  const handleDelete = async (id) => {
    await axios.delete(`/admin/users/${id}`);
    setUsers(users.filter((user) => user.id !== id));
  };

  return (
    <div className="p-10">
      <h1 className="text-2xl font-bold">Admin Dashboard</h1>
      <ul>
        {users.map((user) => (
          <li key={user.id} className="flex justify-between p-2 border-b">
            {user.email} <button onClick={() => handleDelete(user.id)} className="bg-gray-500 text-white px-2">Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AdminDashboard;
