import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./Home.module.css";

function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setIsLoggedIn(true);
    } else {
      setIsLoggedIn(false);
    }
  }, []);

  return (
    <div className={styles.container}>
      {isLoggedIn ? (
        <>
          <div>
            <input type="text" />
            <button />
          </div>
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
