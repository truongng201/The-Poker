import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import "bootstrap/dist/css/bootstrap.min.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import RootLayout from "./components/RootLayout";
import Monitor from "./pages/Monitor";
import Poker from "./pages/Poker";
import Error from "./pages/Error";
import Signin from "./pages/game/auth/Signin";
import Signup from "./pages/game/auth/Signup";
import Home from "./pages/game/Home";
import ForgotPassword from "./pages/game/auth/ForgotPassword";
import ResetPassword from "./pages/game/auth/ResetPassword";
import Join from "./pages/game/Join";
import Create from "./pages/game/Create";

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
        children: [
          { path: "/create", element: <Create />},
          { path: "/join", element: <Join /> },
          { path: "/reset/:resetID", element: <ResetPassword /> },
          { path: "/forgot", element: <ForgotPassword /> },
          { path: "/signup", element: <Signup /> },
          { path: "/signin", element: <Signin /> },
          { path: "/", element: <Home /> },
        ],
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
