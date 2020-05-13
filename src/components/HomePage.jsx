import React from "react";
import { Input } from "antd";
import "../style.css";
import Header from "./Header.jsx";
import Footer from "./Footer.jsx";

const { Search } = Input;

function HomePage() {
  return (
    <>
      <Header />
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
      <Footer />
    </>
  );
}

export default HomePage;
