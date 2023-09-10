import { Outlet, Link } from "react-router-dom";
import "./RootLayout.css";
import React from "react";
import HomeIcon from "../assets/icons/home.png";
import MonitorIcon from "../assets/icons/monitor.png";
import BlogIcon from "../assets/icons/blog.png";
import CardsIcon from "../assets/icons/cards.png";
import Logo from "../assets/logo/logo-no-background.svg";
import Avatar from "../assets/avatar.png";

export default function RootLayout() {
  const pages = [
    {
      name: "Home",
      link: "https://github.com/truongng201/The-Poker/blob/main/README.md",
      className: "page-component home",
      iconPath: HomeIcon,
    },
    {
      name: "Game",
      link: "/",
      className: "page-component game",
      iconPath: CardsIcon,
    },
    {
      name: "Monitor",
      link: "/monitor",
      className: "page-component monitor",
      iconPath: MonitorIcon,
    },
    {
      name: "Blog",
      link: "https://truongng012.notion.site/The-Poker-59eedcb297c340938be0daf79ec7c68f",
      className: "page-component blog",
      iconPath: BlogIcon,
    },
  ];
  return (
    <div className="root-layout">
      <div className="left-container">
        <Link className="logo" to="/">
          <img src={Logo} alt="logo" className="logo" />
        </Link>
        <div className="pages">
          {pages.map((page, index) => {
            return (
              <Link key={index} className={page.className} to={page.link}>
                <img
                  width={16}
                  height={16}
                  src={page.iconPath}
                  alt="page icon"
                />
                <div>{page.name}</div>
              </Link>
            );
          })}
        </div>
        <div className="developer-info">
          <div className="username">Truong Nguyen</div>
          <img src={Avatar} className="avatar" alt="developer avatar" />
          <div className="bio">
            I&apos;m a software engineer who loves to build things. I write
            about building scalable systems.
          </div>
          <div className="social-media">
            <div className="github">
              <a href="https://github.com/truongng201">
                <i className="fab fa-github"></i>
              </a>
            </div>
            <div className="facebook">
              <a href="https://facebook.com">
                <i className="fab fa-facebook"></i>
              </a>
            </div>
            <div className="linkeding">
              <a href="https://linkedin.com">
                <i className="fab fa-linkedin"></i>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div className="right-container">
        <Outlet />
      </div>
    </div>
  );
}
