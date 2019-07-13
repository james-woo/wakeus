import React, { Component }  from 'react';

export class RainbowForm extends Component {
  constructor(props) {
    super(props);
    this.perform = this.perform.bind(this);
  }

  perform() {
    fetch('http://localhost:8000/api/command/rainbow', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Rainbow success:", response.statusText);
        } else {
          console.log("Rainbow error");
        }
      },
      (reason) => {
        console.log("Rainbow error:", reason);
      });
  }

  render() {
    return (
      <div>
        <button style={{padding:10}} onClick={this.perform}>Rainbow</button>
      </div>
    );
  }
}