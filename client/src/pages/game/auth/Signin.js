import "./Signin.css";

import React, { useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

import BackIcon from "../../../assets/icons/back.png";
import AlertComponent from "../../../components/Alerts";
import { EmailValidation, PasswordValidation } from "../../../utils/validation";
import LoadingComponent from "../../../components/Loading";

export default function Signin() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const signin = (e) => {
    e.preventDefault();
    if (!EmailValidation(email)) {
      setErrorMessage("Invalid email address");
    } else if (!PasswordValidation(password)) {
      setErrorMessage(
        "Password must be 8-32 characters long and contain at least one letter and one number"
      );
    } else {
      setErrorMessage("");
      setIsLoading(true);
      axios
        .post(`${process.env.REACT_APP_SERVER_URL}/auth/signin`, {
          email: email,
          password: password,
        })
        .then((res) => {
          localStorage.setItem("access_token", res.data.payload?.access_token);
          setIsLoading(false);
          navigate("/");
        })
        .catch((err) => {
          setErrorMessage(err.response?.data?.message);
          if (err.response?.data?.message === "Email not verified") {
            setErrorMessage(
              "Email not verified yet. We will redirect you to the verification page"
            );
            setTimeout(() => {
              navigate("/reverify");
            }, 5000);
          }
          setIsLoading(false);
        });
    }

    setTimeout(() => {
      setErrorMessage("");
    }, 5000);
  };

  return (
    <div className="Signin">
      <div className="signin-upper-container">
        <div className="back">
          <Link to="/">
            <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
            <span>Back</span>
          </Link>
        </div>
        <form className="shared-form">
          <div className="shared-form-group form-group-signin">
            <input
              type="text"
              className="shared-form-control form-control-signin"
              id="username"
              placeholder="thepoker@gmail.com"
              onChange={(e) => setEmail(e.target.value)}
            />
            <input
              type="password"
              className="shared-form-control form-control-signin"
              id="password"
              placeholder="Password"
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button
            className="signin-button"
            onClick={signin}
            onEnter={signin}
            disabled={isLoading}>
            {isLoading ? (
              <LoadingComponent size={"sm"} />
            ) : (
              <span>Sign In</span>
            )}
          </button>
          <div className="signin-group">
            <Link to="/forgot" className="forgot-password">
              Forgot password ?
            </Link>
            <Link to="/signup" className="signup">
              Sign Up
            </Link>
          </div>
        </form>
      </div>
      <div>
        <div className="or-divider">
          <span className="line"></span>
          <span className="or-text">OR</span>
          <span className="line"></span>
        </div>
        <div className="oauth-groups">
          <div className="oauth-button google-button">
            <a href="google.com">
              <i className="fab fa-google"></i>
            </a>
          </div>
          <div className="oauth-button facebook-button">
            <a href="facebook.com">
              <i className="fab fa-facebook"></i>
            </a>
          </div>
          <div className="oauth-button github-button">
            <a href="github.com">
              <i className="fab fa-github"></i>
            </a>
          </div>
          <div className="oauth-button discord-button">
            <a href="discord.com">
              <i className="fab fa-discord"></i>
            </a>
          </div>
        </div>
      </div>
      {errorMessage !== "" && (
        <AlertComponent
          message={errorMessage}
          variant={"danger"}
          handleClose={() => setErrorMessage("")}
        />
      )}
    </div>
  );
}
