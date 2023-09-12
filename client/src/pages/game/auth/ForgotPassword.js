import React from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link } from "react-router-dom";
import "./ForgotPassword.css";

export default function ForgotPassword() {
  return (
    <div className="shared-container ForgotPassword">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form className="shared-form">
        <div className="shared-form-group form-group-forgotpassword">
          <input
            type="text"
            className="shared-form-control form-control-forgotpassword"
            id="email"
            placeholder="Email"
          />
          <div className="shared-button forgotpassword-button">
            <span>Send</span>
          </div>
        </div>
      </form>
    </div>
  );
}
