import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./Home.module.css";

function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [formData, setFormData] = useState({ name: "" });
  const [message, setMessage] = useState("");
  const [users, setUsers] = useState([]); // State to hold fetched users
  const [suggestions, setSuggestions] = useState([]); // State to hold suggestions
  const navigate = useNavigate();
  const fetcher = "http://192.168.210.89:8080/api/searchUsers";

  useEffect(() => {
    const token = localStorage.getItem("token");
    setIsLoggedIn(!!token);
  }, []);

  // Function to debounce the search
  const debounce = (func, delay) => {
    let timeoutId;
    return (...args) => {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        func(...args);
      }, delay);
    };
  };

  // Function to handle input changes
  const handleChange = (e) => {
    const searchTerm = e.target.value;
    setFormData({ ...formData, name: searchTerm }); // Update the state with the input value
    if (searchTerm === "") {
      setSuggestions([]); // Clear suggestions if input is empty
    } else {
      debouncedSearchChange(searchTerm); // Debounced search
    }
  };

  // Search handler with debounce
  const handleSearchChange = async (searchTerm) => {
    setMessage("Searching...");
    try {
      const response = await fetch(fetcher, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name: searchTerm }),
        credentials: "include",
      });

      if (!response.ok) {
        const data = await response.json();
        setMessage(data.error || "Fetching users failed.");
        return;
      }

      const data = await response.json();
      if (data.users) {
        setSuggestions(data.users); // Set suggestions based on search response
      } else {
        setSuggestions([]); // Clear suggestions if no users
      }
    } catch (error) {
      setMessage("Error fetching data");
      console.error(error);
    }
  };

  // Debounced version of the search change handler
  const debouncedSearchChange = debounce(handleSearchChange, 500);

  return (
    <div className={styles.container}>
      {isLoggedIn ? (
        <>
          <div className={styles.top_left}>
            <form
              className={styles.searchForm}
              onSubmit={(e) => e.preventDefault()} // Prevent form submission
            >
              <input
                type="text"
                name="name"
                placeholder="Search..."
                value={formData.name}
                onChange={handleChange} // Use handleChange to update input value
                className={styles.searchInput}
              />
              <button type="submit" className={styles.searchButton}>
                Search
              </button>
            </form>
            {/* Suggestions dropdown */}
            {suggestions.length > 0 && (
              <div className={styles.suggestions}>
                {suggestions.map((user) => (
                  <div key={user.ID} className={styles.suggestionItem}>
                    <p>{user.Email}</p>
                  </div>
                ))}
              </div>
            )}
          </div>
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
