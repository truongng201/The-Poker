import React from "react";
import "./PokerTool.css";

export default function PokerTool() {
  return (
    <div className="PokerTool">
      <button className="func-button raise-button">Raise</button>
      <button className="func-button call-button">Call</button>
      <button className="func-button fold-button">Fold</button>
    </div>
  );
}
