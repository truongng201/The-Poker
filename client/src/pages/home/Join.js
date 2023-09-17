import "./Join.css";
import React from "react";
import BackIcon from "../../assets/icons/back.png";
import { Link } from "react-router-dom";

export default function Join() {
  return (
    <div className="shared-container Join">
      <div className="back">
        <Link to="/">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <div className="shared-title join-title">Let&apos;s join a room</div>
      <form className="shared-form">
        <div className="shared-form-group" id="form-group-join">
          <input
            type="text"
            className="shared-form-control form-control-join"
            id="roomid"
            placeholder="Room ID"
          />
          <div className="shared-button join-button">
            <span>Join</span>
          </div>
        </div>
      </form>
    </div>
  );
}
