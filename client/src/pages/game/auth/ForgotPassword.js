import "./ForgotPassword.css";
import React from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link } from "react-router-dom";

export default function ForgotPassword() {
  return (
    <div className="ForgotPassword">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form>
        <div className="form-group-forgotpassword">
          <input
            type="text"
            className="form-control-forgotpassword"
            id="email"
            placeholder="Email"
          />
          <div className="forgotpassword-button">
            <span>Send</span>
          </div>
        </div>
      </form>
    </div>
  );
}
