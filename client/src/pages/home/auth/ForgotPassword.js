import "./ForgotPassword.css";
import React, { useState } from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";
import { EmailValidation } from "../../../utils/validation";
import AlertComponent from "../../../components/Alerts";
import LoadingComponent from "../../../components/Loading";

export default function ForgotPassword() {
  const [email, setEmail] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const forgot = (e) => {
    e.preventDefault();
    if (!EmailValidation(email)) {
      setErrorMessage("Invalid email address");
    } else {
      setErrorMessage("");
      setIsLoading(true);
      axios
        .post(`${process.env.REACT_APP_SERVER_URL}/auth/forgot-password`, {
          email: email,
        })
        .then(() => {
          setIsLoading(false);
          navigate("/reset-email");
        })
        .catch((err) => {
          setErrorMessage(err.response?.data?.message);
          setIsLoading(false);
        });
    }
  };

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
            placeholder="thepoker@gmail.com"
            onChange={(e) => {
              setEmail(e.target.value);
            }}
          />
          <button
            className="shared-button forgotpassword-button"
            disabled={isLoading}
            onClick={forgot}>
            {isLoading ? <LoadingComponent size="sm" /> : <span>Send</span>}
          </button>
        </div>
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
