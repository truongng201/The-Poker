import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import "bootstrap/dist/css/bootstrap.min.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import RootLayout from "./components/RootLayout";
import Monitor from "./pages/Monitor";
import Poker from "./pages/Poker";
import Error from "./pages/Error";
import Login from "./pages/poker/Login";

const router = createBrowserRouter([
  {
    path: "/",
    element: window.innerWidth >= 1200 ? <RootLayout /> : <Error />,
    errorElement: <Error />,
    children: [
      { path: "/monitor", element: <Monitor /> },
      {
        path: "/",
        element: <Poker />,
        children: [{ path: "/", element: <Login /> }],
      },
    ],
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
