import { Link } from "react-router-dom";
import styles from "./NavBar.module.css"; // Import the CSS module

function NavBar() {
  return (
    <div className={styles.navbar}>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/login">Login</Link>
        </li>
        <li>
          <Link to="/register">Register</Link>
        </li>
        <li>
          <Link to="/logout">Logout</Link>
        </li>
      </ul>
    </div>
  );
}

export default NavBar;
