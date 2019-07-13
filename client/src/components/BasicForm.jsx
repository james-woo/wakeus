import React, { Component }  from 'react';
import { ChromePicker } from 'react-color';

export class BasicForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      color: '#fff'
    };

    this.handleChange = this.handleChange.bind(this);
    this.perform = this.perform.bind(this);
  }

  handleChange(color, event) {
    this.setState({
      color: color.rgb
    });
    console.log(color.rgb);
  }

  perform() {
    console.log(`Request basic: 
      color: ${this.state.color},  
      intensity: ${this.state.color.a}`
    );
    let host = process.env.API_SERVICE_HOST || "localhost";
    fetch(`http://${host}:8000/api/command/basic`, {
      method: 'POST',
      cors: 'enabled',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        color: {
          r: this.state.color.r,
          g: this.state.color.g,
          b: this.state.color.b
        },
        intensity: this.state.color.a,
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
  }

  render() {
    return (
      <div>
        <ChromePicker
          color={this.state.color}
          onChange={this.handleChange}
        />
        <button style={{padding:10}} onClick={this.perform}>Basic</button>
      </div>
    );
  }
}