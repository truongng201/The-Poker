import "./Home.css";
import { Link } from "react-router-dom";
import DoorIcon from "../../assets/icons/door.png";
import GameControllerIcon from "../../assets/icons/game-controller.png";
import React, { useEffect, useState } from "react";

export default function Home() {
  const [username, setUsername] = useState("");
  useEffect(() => {
    if (
      sessionStorage.getItem("username") === null ||
      sessionStorage.getItem("username") === ""
    ) {
      const username = `Alpha${Math.floor(Math.random() * 100000)}`;
      sessionStorage.setItem("username", username);
    }
    setUsername(sessionStorage.getItem("username"));
  }, []);

  return (
    <div className="Home-Poker">
      <div className="button-groups">
        <Link to="/join">
          <div className="home-button join-button">
            <img
              className="login-join-icon"
              alt="icon"
              src={DoorIcon}
              width={30}
              height={30}
            />
            <span>Join</span>
          </div>
        </Link>

        <Link to="/create">
          <div className="home-button create-button">
            <img
              className="login-create-icon"
              alt="icon"
              src={GameControllerIcon}
              width={30}
              height={30}
            />
            <span>Create</span>
          </div>
        </Link>
      </div>
      <div>
        <div className="home-title">Join to play with us as</div>
        <div className="login-input">
          <div className="login-avatar-block">
            <img
              className="login-avatar"
              alt="rand-avt"
              src={`https://api.dicebear.com/6.x/bottts-neutral/svg?seed=${username}`}
              width={80}
              height={80}
            />
          </div>
          <div className="login-input-rand">
            <input
              type="text"
              placeholder="Username"
              defaultValue={username}
              onChange={(e) => {
                setUsername(e.target.value);
                sessionStorage.setItem("username", e.target.value);
              }}
            />
          </div>
        </div>
        <div className="or-divider">
          <span className="line"></span>
          <span className="or-text">OR</span>
          <span className="line"></span>
        </div>
        <div className="login-home">
          <Link to="/signin" className="continue-login-button">
            <span>Continue to sign in</span>
          </Link>
        </div>
      </div>
    </div>
  );
}
