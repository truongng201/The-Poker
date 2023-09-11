import "./Join.css";
import React from "react";
import BackIcon from "../../assets/icons/back.png";
import { Link } from "react-router-dom";

export default function Join() {
  return (
    <div className="Join">
      <div className="back">
        <Link to="/">
          <img src={BackIcon} alt="icon" width={"16px"} height={"16px"} />
          <span>Back</span>
        </Link>
      </div>
      <div className="join-title">Let&apos;s join a room</div>
      <form>
        <div className="form-group-join">
          <input
            type="text"
            className="form-control-join"
            id="roomid"
            placeholder="Room ID"
          />
          <div className="join-button">
            <span>Join</span>
          </div>
        </div>
      </form>
    </div>
  );
}
