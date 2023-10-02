import React from "react";
import { useNavigate, useParams } from "react-router-dom";
import "./ResetPassword.css";
import axios from "axios";
import { PasswordValidation } from "../../../utils/validation";
import AlertComponent from "../../../components/Alerts";
import LoadingComponent from "../../../components/Loading";

export default function ResetPassword() {
  const [isLoading, setIsLoading] = React.useState(false);
  const [errorMessage, setErrorMessage] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [retypePassword, setRetypePassword] = React.useState("");
  const navigate = useNavigate();

  const { resetID } = useParams();

  const reset = (e) => {
    e.preventDefault();
    if (!PasswordValidation(password)) {
      setErrorMessage(
        "Password must be 8-32 characters long and contain at least one letter and one number"
      );
    } else if (password !== retypePassword) {
      setErrorMessage("Passwords do not match");
    } else {
      setErrorMessage("");
      setIsLoading(true);
      axios
        .post(`${process.env.REACT_APP_SERVER_URL}/auth/reset-password`, {
          reset_token: resetID,
          new_password: password,
        })
        .then(() => {
          setIsLoading(false);
          navigate("/confirm-reset");
        })
        .catch((err) => {
          setErrorMessage(err.response?.data?.message);
          setIsLoading(false);
        });
    }
  };

  return (
    <div className="shared-container ResetPassword">
      <div className="shared-title">Reset your password</div>
      <form className="shared-form">
        <div className="shared-form-group form-group-resetpassword">
          <input
            type="password"
            className="shared-form-control form-control-resetpassword"
            id="newpassword"
            placeholder="New Password"
            onChange={(e) => {
              setPassword(e.target.value);
            }}
          />
          <input
            type="password"
            className="shared-form-control form-control-resetpassword"
            id="retype-newpassword"
            placeholder="Retype New Password"
            onChange={(e) => {
              setRetypePassword(e.target.value);
            }}
          />
        </div>
        <button
          className="shared-button"
          id="resetpassword-button"
          onClick={reset}
          disabled={isLoading}>
          {isLoading ? <LoadingComponent size={"sm"} /> : <span>Reset</span>}
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
