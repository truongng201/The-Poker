import "./NotFoundPage.css";
import React from "react";

export default function NotFoundPage() {
  return (
    <div className="NotFoundPage">
      <div className="NotFoundMessage">
        <h1>404</h1>
        <p>
          This page could not be found or you are using an unsupported device.
          Please open this website on a desktop or laptop computer. We apologize
          for the inconvenience.
        </p>
      </div>
    </div>
  );
}
