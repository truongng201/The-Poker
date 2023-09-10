import "./Login.css";
import React from "react";
import { Button } from "react-bootstrap";

export default function Login() {
  return (
    <div className="Login">
      <div className="login-upper-container">
        <div className="login-title">The Poker</div>
        <form>
          <div className="form-group">
            <input
              type="text"
              className="form-control"
              id="username"
              placeholder="Email"
            />
            <input
              type="password"
              className="form-control"
              id="password"
              placeholder="Password"
            />
          </div>
          <Button variant="primary" type="submit">
            Login
          </Button>
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
