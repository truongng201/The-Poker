import React from "react";
import BackIcon from "../../../assets/icons/back.png";
import { Link, useParams } from "react-router-dom";

export default function ResetPassword() {
  const { resetID } = useParams();
  console.log(resetID);
  return (
    <div className="shared-container ResetPassword">
      <div className="back">
        <Link to="/signin">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <form className="shared-form">
        <div className="shared-form-group form-group-resetpassword">
          <input
            type="text"
            className="shared-form-control form-control-resetpassword"
            id="newpassword"
            placeholder="New Password"
          />
          <input
            type="text"
            className="shared-form-control form-control-resetpassword"
            id="retype-newpassword"
            placeholder="Retype New Password"
          />
        </div>
        <div className="shared-button resetpassword-button">
          <span>Reset</span>
        </div>
      </form>
    </div>
  );
}
