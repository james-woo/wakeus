import React, {Component} from 'react';
import {ColorPicker} from "./ColorPicker";
import {Button, Segment} from "semantic-ui-react";

export class BasicForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      color: {
        rgb: {r: 250, g: 0, b: 4, a: 1.0},
        hsl: {h: 360, s: 1, l:0.5, a:1.0}
      }
    };

    this.onChange = this.onChange.bind(this);
    this.perform = this.perform.bind(this);
  }

  onChange(color) {
    this.setState(color);
  }

  perform() {
    let host = process.env.REACT_APP_HOST_IP_ADDRESS || "192.168.1.52";
    console.log(`Request basic: 
      color: ${JSON.stringify(this.state.color.rgb)},
      intensity: ${this.state.color.rgb.a} 
      to host ${host}`
    );
    fetch(`http://${host}:8000/api/command/basic`, {
      method: 'POST',
      cors: 'enabled',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        color: {
          r: this.state.color.rgb.r,
          g: this.state.color.rgb.g,
          b: this.state.color.rgb.b
        },
        intensity: this.state.color.rgb.a,
      })
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Basic success:", response.statusText);
        } else {
          console.log("Basic error");
        }
      },
      (reason) => {
        console.log("Basic error:", reason);
      });
  };

  render() {
    return (
      <Segment inverted>
        <h1>Basic</h1>
        <ColorPicker
          color={this.state.color}
          onChange={this.onChange}/>
        <br/>
        <Button onClick={this.perform}>Perform</Button>
      </Segment>
    );
  }
}