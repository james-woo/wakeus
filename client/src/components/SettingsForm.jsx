import React, {Component} from 'react';
import {Segment, Checkbox} from "semantic-ui-react";

export class SettingsForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      pause: false,
    };
    this.handlePause = this.handlePause.bind(this);
  }

  componentDidMount() {
    let host = process.env.REACT_APP_HOST_IP_ADDRESS || "192.168.1.52";
    fetch(`http://${host}:8000/api/settings`, {
        method: 'GET'
      }).then(
        (response) => {
          if (response.ok) {
            console.log("Setting success:", response.statusText);
          } else {
            console.log("Setting error");
          }
          response.json();
        },
        (reason) => {
          console.log("Setting error:", reason);
        }).then(data => this.setState({ data }));
  }

  handlePause(event) {
    this.setState({
      pause: !this.state.pause
    });

    let host = process.env.REACT_APP_HOST_IP_ADDRESS || "192.168.1.52";
    console.log(`Update settings: ${this.state.pause}`);
    fetch(`http://${host}:8000/api/settings`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        pause: this.state.pause
      })
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Setting success:", response.statusText);
        } else {
          console.log("Setting error");
        }
      },
      (reason) => {
        console.log("Setting error:", reason);
      });
  }

  render() {
    return (
      <Segment>
        <h1>Settings</h1>
        <Checkbox
          label='Pause'
          toggle
          onChange={this.handlePause}/>
      </Segment>
    );
  }
}