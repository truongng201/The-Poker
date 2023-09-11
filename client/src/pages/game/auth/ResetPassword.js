import "./ResetPassword.css";
import React from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link, useParams } from "react-router-dom";

export default function ResetPassword() {
  const { resetID } = useParams();
  console.log(resetID);
  return (
    <div className="ResetPassword">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form>
        <div className="form-group-resetpassword">
          <input
            type="text"
            className="form-control-resetpassword"
            id="newpassword"
            placeholder="New Password"
          />
          <input
            type="text"
            className="form-control-resetpassword"
            id="retype-newpassword"
            placeholder="Retype New Password"
          />
        </div>
        <div className="resetpassword-button">
          <span>Reset</span>
        </div>
      </form>
    </div>
  );
}
