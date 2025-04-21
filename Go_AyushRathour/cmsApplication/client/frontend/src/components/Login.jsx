import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "../api/axios";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("/login", { email, password });
      localStorage.setItem("token", res.data.token);
      navigate("/user/dashboard");
    } catch (err) {
      console.error("Login failed", err);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <img src="/logo.png" alt="Logo" className="mb-4 w-20 h-20" />
      <form onSubmit={handleLogin} className="flex flex-col space-y-4">
        <input className="border p-2" type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
        <input className="border p-2" type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
        <button type="submit" className="bg-gray-500 text-white p-2">Login</button>
      </form>
    </div>
  );
};

export default Login;
