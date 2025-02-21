import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function Logout() {
  const navigate = useNavigate();

  useEffect(() => {
    // Clear JWT token from localStorage
    localStorage.removeItem("token");

    // Inform the server about logout (optional, not needed if only using JWT on frontend)
    const fetcher = "http://192.168.210.89:8080/api/logout";
    fetch(fetcher, {
      method: "POST",
      credentials: "include", // Only needed if still using cookies
    });

    // Redirect to login page after logout
    navigate("/login");
  }, [navigate]);

  return <h1>Logout Successful</h1>;
}

export default Logout;
