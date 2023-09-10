import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import RootLayout from "./components/RootLayout";
import Monitor from "./pages/Monitor";
import Home from "./pages/Home";
import Poker from "./pages/Poker";
import Blog from "./pages/Blog";
import Error from "./pages/Error";

const router = createBrowserRouter([
  {
    path: "/",
    element: window.innerWidth >= 1200 ? <RootLayout /> : <Error />,
    errorElement: <Error />,
    children: [
      { path: "/", element: <Home /> },
      { path: "/monitor", element: <Monitor /> },
      { path: "/poker", element: <Poker /> },
      { path: "/blog", element: <Blog /> },
    ],
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
