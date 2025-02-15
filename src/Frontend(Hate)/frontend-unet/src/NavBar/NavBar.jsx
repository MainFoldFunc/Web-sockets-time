// NavBar.js
import styles from "./NavBar.module.css"; // Import the CSS module

function NavBar() {
  return (
    <div className={styles.navbar}>
      {" "}
      {/* Apply the navbar class */}
      <ul>
        <li>
          <a href="#home">Home</a>
        </li>
        <li>
          <a href="#about">Login</a>
        </li>
        <li>
          <a href="#services">Register</a>
        </li>
        <li>
          <a href="#contact">register</a>
        </li>
      </ul>
    </div>
  );
}

export default NavBar;
