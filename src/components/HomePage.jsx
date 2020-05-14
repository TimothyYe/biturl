import React, { useState } from "react";
import { Input, Button } from "antd";
import "../style.css";
import Header from "./Header.jsx";
import Footer from "./Footer.jsx";

function HomePage() {
  const [urlValue, setURL] = useState("");

  const onShorten = (value, e) => {
    console.log(value);
    setURL("*******");
    console.log(e.target);
    e.target.type = "dashed";
  };

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
          width: "25vw",
          textAlign: "center",
        }}
      >
        <h1>BitURL</h1>
        <h3>Make your links manageable.</h3>
        <br />

        <div className="input-group">
          <div className="input-box">
            <Input
              readOnly={false}
              placeholder="Please enter URL..."
              size="default"
            />
          </div>
          <Button type="primary" size="default">
            SHORTEN
          </Button>
        </div>
      </div>
      <Footer />
    </>
  );
}

export default HomePage;
