import React from "react";
import { Input } from "antd";
import "../style.css";

const { Search } = Input;

function HomePage() {
  return (
    <div
      style={{
        position: "absolute",
        left: "50%",
        top: "50%",
        transform: "translate(-50%, -50%)",
        color: "white",
      }}
    >
      <h1>BitURL</h1>
      <h3>Make your links manageable.</h3>
      <br />

      <div>
        <Search
          placeholder="Please enter URL..."
          enterButton="SHORTEN"
          size="large"
          onSearch={(value) => console.log(value)}
        />
      </div>
    </div>
  );
}

export default HomePage;
