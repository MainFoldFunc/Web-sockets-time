import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./Home.module.css";

function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [formData, setFormData] = useState({ name: "" });
  const [message, setMessage] = useState("");
  const [users, setUsers] = useState([]); // State to hold fetched users
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    setIsLoggedIn(!!token);
  }, []);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault(); // Fixed typo here (preventDefault instead of PreventDefault)

    if (formData.name === "") {
      setMessage("Please fill in the field");
      return;
    }

    setMessage("Searching for users");

    try {
      const response = await fetch("http://localhost:8080/api/search", {
        // Changed to a hypothetical search endpoint
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
        credentials: "include",
      });

      const data = await response.json();

      if (response.ok) {
        setMessage("Users found!");
        setUsers(data.users); // Assuming the response contains users in the "users" property
      } else {
        setMessage(
          data.error || "Fetching users failed, please try again later",
        );
      }
    } catch (error) {
      setMessage("Network error. Please try again later.");
    }
  };

  return (
    <div className={styles.container}>
      {isLoggedIn ? (
        <>
          <div className={styles.top_left}>
            <input
              type="text"
              name="name"
              placeholder="Search..."
              value={formData.name}
              onChange={handleChange}
            />
            <button type="submit" onClick={handleSubmit}>
              Search
            </button>
          </div>
          {message && <p>{message}</p>} {/* Display message */}
          {users.length > 0 && (
            <div className={styles.userList}>
              {users.map((user) => (
                <div key={user.id} className={styles.userCard}>
                  <h3>{user.name}</h3>
                  <p>{user.email}</p>
                </div>
              ))}
            </div>
          )}
        </>
      ) : (
        <>
          <h1>Welcome to the Main Page</h1>
          <p>Please log in to access more features.</p>
          <button onClick={() => navigate("/login")}>Login</button>
        </>
      )}
    </div>
  );
}

export default Home;
