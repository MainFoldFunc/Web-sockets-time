import styles from "./Login.module.css";

function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const [message, setMessage] = useState("");

  const handleChange = (e) => {
    setFormData({ ...fromData, [e.target.name]: e.tearget.value });
  };
  return (
    <div className={styles.container}>
      <form>
        <h2>Login</h2>
        <input type="email" placeholder="Email" required />
        <input type="password" placeholder="Password" required />
        <button type="submit">Sign Up</button>
      </form>
    </div>
  );
}

export default Login;
