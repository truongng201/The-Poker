import React from "react";
import BackIcon from "../../assets/icons/back.png";
import { Link } from "react-router-dom";

export default function Create() {
  return (
    <div className="shared-container Create">
      <div className="back">
        <Link to="/">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <div className="shared-title create-title">
        Let&apos;s create new room
      </div>
      <form className="shared-form">
        <div className="shared-form-group form-group-create">
          <div className="shared-button create-button">
            <span>Create</span>
          </div>
        </div>
      </form>
    </div>
  );
}
