import "./SentEmail.css";

import React from "react";
import { Link } from "react-router-dom";

import BackIcon from "../../../assets/icons/back.png";

export default function SentEmail({ email_type }) {
  const reset_content = `
    We have sent you an email to reset your password. 
    Please check your email and click on the link to reset your password.
    If you did not receive the email, please check your spam folder.
  `;

  const verify_content = `
    We have sent you an email to verify your account. 
    Please check your email and click on the link to verify your account.
    If you did not receive the email, please check your spam folder.
  `;
  return (
    <div className="shared-container">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <div className="sent-email-content">
        {email_type === "reset" && reset_content}
        {email_type === "reset" && (
          <Link to="/forgot" className="resend">
            Reset here
          </Link>
        )}
        {email_type === "verify" && verify_content}
        {email_type === "verify" && (
          <Link to="/reverify" className="resend">
            Reverify here
          </Link>
        )}
      </div>
    </div>
  );
}
