import React from "react";
import "./Room.css";
import TablePoker from "../../components/room/TablePoker";
import Chatbox from "../../components/room/Chatbox";
import PokerTool from "../../components/room/PokerTool";

export default function Room() {
  return (
    <div className="RoomContainer">
      <div className="Room">
        <div className="left-room-container">
          <TablePoker />
          <PokerTool />
        </div>
        <Chatbox />
      </div>
    </div>
  );
}
