import "./ConfirmEmail.css";

import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";

import SuccessIcon from "../../../assets/icons/correct.png";

export default function ConfirmEmail({ email_type }) {
  const navigate = useNavigate();
  const reset_content = `
    You have successfully reset your password. 
    Please wait for few seconds to be redirected to the login page.
  `;

  const verify_content = `
    You have successfully verified your account. 
    Please wait for few seconds to be redirected to the login page.
  `;

  useEffect(() => {
    setTimeout(() => {
      navigate("/signin");
    }, 10000);
  });

  return (
    <div className="shared-container">
      <img src={SuccessIcon} alt="icon" width={"50px"} height={"50px"} />
      <div className="confirm-email-content">
        {email_type === "reset" && reset_content}
        {email_type === "verify" && verify_content}
      </div>
    </div>
  );
}
