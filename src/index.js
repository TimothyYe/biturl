import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router } from "react-router-dom";
import MainContainer from "./components/MainContainer.jsx";

ReactDOM.render(
  <Router>
    <MainContainer />
  </Router>,
  document.getElementById("main-container")
);
