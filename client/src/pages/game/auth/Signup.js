import "./Signup.css";
import React from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link } from "react-router-dom";

export default function Signup() {
  return (
    <div className="Signup">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form>
        <div className="form-group-signup">
          <input
            type="text"
            className="form-control-signup"
            id="email"
            placeholder="Email"
          />
          <input
            type="text"
            className="form-control-signup"
            id="username"
            placeholder="Username"
          />
          <input
            type="password"
            className="form-control-signup"
            id="password"
            placeholder="Password"
          />
          <input
            type="retype-password"
            className="form-control-signup"
            id="retype-password"
            placeholder="Retype Password"
          />
        </div>
        <div className="signup-button">
          <span>Sign Up</span>
        </div>
      </form>
    </div>
  );
}
