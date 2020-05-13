import React from "react";
import "../style.css";

function Header() {
  return (
    <div
      className="header"
      style={{
        position: "absolute",
        top: "5%",
        right: "5%",
        color: "white",
        height: "2vh",
        transform: "translate(-50%, -50%)",
      }}
    >
      <ul>
        <li>Home</li>
        <li>About</li>
      </ul>
    </div>
  );
}

export default Header;
