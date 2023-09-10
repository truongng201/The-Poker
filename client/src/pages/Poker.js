import "./Poker.css";
import React from "react";
import Logo from "../assets/logo/logo-no-background.svg";
import { Outlet } from "react-router-dom";

export default function Poker() {
  return (
    <div className="Poker">
      <div className="poker-layout">
        <div className="poker-left-container">
          <Outlet />
        </div>
        <div className="poker-right-container">
          <img
            className="login-logo"
            alt="logo"
            src={Logo}
            width={300}
            height={100}
          />
        </div>
      </div>
    </div>
  );
}
