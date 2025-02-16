import { useState } from "react";
import { useNavigate } from "react-router-dom"; // ✅ Import useNavigate
import styles from "./Login.module.css";

function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const [message, setMessage] = useState("");
  const navigate = useNavigate(); // ✅ Define navigate

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (formData.email === "" || formData.password === "") {
      setMessage("Please fill in all fields.");
      return;
    }
    setMessage("Logging in...");

    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
        credentials: "include",
      });

      const data = await response.json();
      if (response.ok) {
        setMessage("Login successful!");
        localStorage.setItem("token", data.token); // ✅ Store JWT token
        setTimeout(() => navigate("/"), 1000); // ✅ Redirect to home after 1 second
      } else {
        setMessage(data.error || "Login failed, please try again later.");
      }
    } catch (error) {
      setMessage("Network error. Please try again later.");
    }
  };

  return (
    <div className={styles.container}>
      <form onSubmit={handleSubmit}>
        <h2>Login</h2>
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          required
        />
        <button type="submit">Login</button>
        {message && <p className={styles.message}>{message}</p>}
      </form>
    </div>
  );
}

export default Login;
