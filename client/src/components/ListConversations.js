import React, { Component } from "react";
import Table from "react-bootstrap/Table";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

import "./ListConversations.css";

class ListConversations extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      conversations: []
    };
  }

  doSearch() {
    fetch("/conversations/author/1", {
      method: "GET",
      headers: {
        Accept: "application/json"
      }
    })
      .then(res => res.json())
      .then(
        result => {
          this.setState({
            isLoaded: true,
            conversations: result.conversations
          });
        }
      );
  }

  render() {
    const {isLoaded, conversations } = this.state;
     if (!isLoaded) {
      return (
        <Form>
          <Form.Group controlId="fromGroupSearchByAuthor">
            <Form.Row>
              <Form.Control
                className="search-input"
                type="text"
                placeholder="Enter author id here"
              />
              <Button
                onClick={this.doSearch.bind(this)}
                className="btn-sm search-button"
                type="submit"
              >
                Search
              </Button>
            </Form.Row>
          </Form.Group>
        </Form>
      );
    } else if (isLoaded && conversations.length > 0) {
      return (
        <div>
          <Form>
            <Form.Group controlId="fromGroupSearchByAuthor">
              <Form.Row>
                <Form.Control
                  className="search-input"
                  type="text"
                  placeholder="Enter author id here"
                />
                <Button
                  onClick={this.doSearch.bind(this)}
                  className="btn-sm search-button"
                  type="submit"
                >
                  Search
                </Button>
              </Form.Row>
            </Form.Group>
          </Form>
          <Table striped bordered hover>
            <thead>
              <tr>
                <th>Conversation ID</th>
                <th>Author ID</th>
                <th>Created At</th>
                <th>Open Conversation</th>
              </tr>
            </thead>
            <tbody>
              {conversations.map(conversation => (
                <tr key={conversation.conversation_id}>
                  <td>{conversation.conversation_id}</td>
                  <td>{conversation.author_id}</td>
                  <td>{conversation.created_at}</td>
                  <td>Open</td>
                </tr>
              ))}
            </tbody>
          </Table>
        </div>
      );
    } else {
      return (
        <div>
          <Form>
            <Form.Group controlId="fromGroupSearchByAuthor">
              <Form.Row>
                <Form.Control
                  className="search-input"
                  type="text"
                  placeholder="Enter author id here"
                />
                <Button
                  onClick={this.doSearch.bind(this)}
                  className="btn-sm search-button"
                  type="submit"
                >
                  Search
                </Button>
              </Form.Row>
            </Form.Group>
          </Form>
          <p>NO RESULTS FOUND</p>
        </div>
      );
    }
  }
}
export default ListConversations;
