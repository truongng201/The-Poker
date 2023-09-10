import "./Monitor.css";
import React from "react";

export default function Monitor() {
  return (
    <div className="Monitor">
      <iframe
        src={process.env.REACT_APP_MONITOR_DASHBOARD_SRC}
        title="monitor"
        width="100%"
        height="100%"
      />
    </div>
  );
}
