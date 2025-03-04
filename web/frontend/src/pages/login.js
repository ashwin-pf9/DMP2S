import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import API_BASE_URL from "../config";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const [errorMessage, setErrorMessage] = useState(""); // State for error messages

  const handleLogin = async (e) => {
    e.preventDefault();
    setErrorMessage(""); // Reset error message before new attempt

    try {
    // REQUEST TO BACKEND 
      const response = await fetch(`${API_BASE_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (response.status === 401) {
        setErrorMessage("Please verify your email before logging in.");
        return;
      }

      if (!response.ok) throw new Error("Login failed");

      const data = await response.json();
      
      if (data.token) {
        localStorage.setItem("token", data.token);
        console.log("Navigating to dashboard...");
        navigate("/dashboard"); // Ensure this runs only once
      }
    } catch (error) {
      console.error("Login error:", error);
      setErrorMessage("Invalid email or password. Please try again.");
    }
  };

  return (
    <div className="container">
    <h2>Login</h2>
    {errorMessage && <p className="error-message">{errorMessage}</p>}
    <form onSubmit={handleLogin}>
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
      <button type="submit">Login</button>
    </form>
    <p>Don't have an account? <a href="/register">Register</a></p>
  </div>
  );
};

export default Login;
