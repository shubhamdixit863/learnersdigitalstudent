import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "../api/axios";

const Signup = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSignup = async (e) => {
    e.preventDefault();
    try {
      await axios.post("/signup", { email, password });
      navigate("/");
    } catch (err) {
      console.error("Signup failed", err);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <img src="/logo.png" alt="Logo" className="mb-4 w-20 h-20" />
      <form onSubmit={handleSignup} className="flex flex-col space-y-4">
        <input className="border p-2" type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
        <input className="border p-2" type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
        <button type="submit" className="bg-gray-500 text-white p-2">Signup</button>
      </form>
    </div>
  );
};

export default Signup;
