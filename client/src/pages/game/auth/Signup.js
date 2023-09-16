import "./Signup.css";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

import BackIcon from "../../../assets/icons/back.png";
import LoadingComponent from "../../../components/Loading";
import {
  EmailValidation,
  PasswordValidation,
  UsernameValidation,
} from "../../../utils/validation";
import AlertComponent from "../../../components/Alerts";

export default function Signup() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [retypePassword, setRetypePassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const signup = (e) => {
    e.preventDefault();
    if (!EmailValidation(email)) {
      setErrorMessage("Invalid email address");
    } else if (!UsernameValidation(username)) {
      setErrorMessage(
        "Username must be 8-32 characters long and contain at least one letter and one number"
      );
    } else if (!PasswordValidation(password)) {
      setErrorMessage(
        "Password must be 8-32 characters long and contain at least one letter and one number"
      );
    } else if (password !== retypePassword) {
      setErrorMessage("Passwords do not match");
    } else {
      setErrorMessage("");
      setIsLoading(true);
      axios
        .post(`${process.env.REACT_APP_SERVER_URL}/auth/signup`, {
          email: email,
          password: password,
          username: username,
        })
        .then(() => {
          localStorage.setItem("email", email);
          setIsLoading(false);
          navigate("/verify-email");
        })
        .catch((err) => {
          setErrorMessage(err.response?.data?.message);
          setIsLoading(false);
        });
    }

    setTimeout(() => {
      setErrorMessage("");
    }, 5000);
  };

  return (
    <div className="shared-container Signup">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form className="shared-form">
        <div className="shared-form-group form-group-signup">
          <input
            type="text"
            className="shared-form-control form-control-signup"
            id="email"
            placeholder="thepoker@gmail.com"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="text"
            className="shared-form-control form-control-signup"
            id="username"
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            className="shared-form-control form-control-signup"
            id="password"
            placeholder="Password"
            onChange={(e) => setPassword(e.target.value)}
          />
          <input
            type="password"
            className="shared-form-control form-control-signup"
            id="retype-password"
            placeholder="Retype Password"
            onChange={(e) => setRetypePassword(e.target.value)}
          />
        </div>
        <button
          className="shared-button"
          id="signup-button"
          onClick={signup}
          onEnter={signup}
          disabled={isLoading}>
          {isLoading ? <LoadingComponent size={"sm"} /> : <span>Sign Up</span>}
        </button>
      </form>
      {errorMessage !== "" && (
        <AlertComponent
          message={errorMessage}
          handleClose={() => setErrorMessage("")}
        />
      )}
    </div>
  );
}
