import React, { useState } from "react";
import "./Card.css";

export default function Card() {
  const [isDown, setIsDown] = useState(true);

  return (
    <div className="card">
      <div className="card-back">
        {isDown && <div className="card-back-img">TP</div>}
      </div>
    </div>
  );
}
