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
    e.preventDefault();

    if (formData.name === "") {
      setMessage("Please fill in the field");
      return;
    }

    setMessage("Searching for users...");

    try {
      const response = await fetch("http://192.168.1.19:8080/api/searchUsers", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
        credentials: "include",
      });

      // Check response.ok before parsing
      if (!response.ok) {
        const data = await response.json();
        setMessage(
          data.error || "Fetching users failed, please try again later.",
        );
        return;
      }

      const data = await response.json();
      console.log(data); // Log the entire response

      if (data.users && data.users.length > 0) {
        setMessage("Users found!");
        setUsers(data.users); // Set the users only if they exist
      } else {
        setMessage("No users found.");
      }
    } catch (error) {
      setMessage("Network error. Please try again later.");
      console.error(error); // Log network errors
    }
  };

  return (
    <div className={styles.container}>
      {isLoggedIn ? (
        <>
          <div className={styles.top_left}>
            <form onSubmit={handleSubmit}>
              <input
                type="text"
                name="name"
                placeholder="Search..."
                value={formData.name}
                onChange={handleChange}
              />
              <button type="submit">Search</button>
            </form>
          </div>{" "}
          {/* Scrollable user list */}
          {users.length > 0 && (
            <div className={styles.userList}>
              {users.map((user) => (
                <div key={user.ID} className={styles.userCard}>
                  <p>{user.Email}</p>
                  <button>Chat</button>
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
