import React from "react";

function Footer() {
  return (
    <div
      className="footer"
      style={{
        position: "absolute",
        bottom: "10%",
        left: "50%",
        color: "white",
        height: "2vh",
        transform: "translate(-50%, -50%)",
      }}
    >
      <h3>
        <strong>BitURL</strong>
      </h3>
      <hr />
      <p>Copyright &copy; BitURL.top 2020</p>
    </div>
  );
}

export default Footer;
