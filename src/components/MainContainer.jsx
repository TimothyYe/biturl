import React, { useState, useEffect } from "react";
import { Route, Switch, Redirect, withRouter } from "react-router";
import { useHistory } from "react-router-dom";
import HomePage from "./HomePage.jsx";
import { getLastUpdate } from "../lib/api.js";
import "../style.css";

function MainContainer() {
  return (
    <div
      style={{
        background: "#fff",
        width: "100vw",
        height: "100vh",
        backgroundImage:
          "url(https://bing.biturl.top/?resolution=1920&format=image&index=0&mkt=zh-CN)",
      }}
    >
      <Switch>
        {/* Home tab */}
        <Route path="/">
          <HomePage style={{ width: "100vw", height: "100vh" }} />
        </Route>
      </Switch>
    </div>
  );
}

export default withRouter(MainContainer);
