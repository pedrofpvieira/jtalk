import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import ListConversations from "./components/ListConversations";
import { Route, Link, HashRouter } from "react-router-dom";

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>JTalk UI - Test Client for Communication Module</h2>
        </div>
        <HashRouter>
          <Navbar bg="light" expand="lg">
            <Navbar.Brand href="/">Home</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="mr-auto">
                <Link
                  className="nav-link"
                  activeclassname="nav-link"
                  to="/list-conversations"
                >
                  List Conversations
                </Link>
                <Link
                  className="nav-link"
                  activeclassame="nav-link"
                  to="/create-conversation"
                >
                  Create Conversation
                </Link>
              </Nav>
            </Navbar.Collapse>
          </Navbar>
          <div className="content">
            <Route path="/list-conversations" component={ListConversations} />
            <Route path="/create-conversation" component={ListConversations} />
          </div>
        </HashRouter>
      </div>
    );
  }
}

export default App;
