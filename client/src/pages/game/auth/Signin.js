import { Link } from "react-router-dom";
import "./Signin.css";
import React from "react";
import BackIcon from "../../../assets/icons/back.png";

export default function Signin() {
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
              placeholder="Email"
            />
            <input
              type="password"
              className="shared-form-control form-control-signin"
              id="password"
              placeholder="Password"
            />
          </div>
          <div className="signin-button">
            <span>Sign In</span>
          </div>
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
    </div>
  );
}
