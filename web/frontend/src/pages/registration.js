import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import API_BASE_URL from "../config";

const Register = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [roleID, setRoleID] = useState(2); // Default role_id (modify if needed)
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${API_BASE_URL}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password, name, role_id: roleID }),
      });

      if (!response.ok) throw new Error("Registration failed");

      navigate("/login"); // Redirect to login after successful registration
    } catch (error) {
      console.error("Registration error:", error);
    }
  };

  return (
    <div className="container">
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <input
          type="text"
          placeholder="Full Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <select value={roleID} onChange={(e) => setRoleID(Number(e.target.value))}>
          <option value={1}>Admin</option>
          <option value={2}>User</option>
        </select>
        <button type="submit">Register</button>
      </form>
      <p>Already have an account? <a href="/login">Login</a></p>
    </div>
  );
};

export default Register;
