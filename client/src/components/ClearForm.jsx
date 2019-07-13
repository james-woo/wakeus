import React, { Component }  from 'react';

export class ClearForm extends Component {
  constructor(props) {
    super(props);
    this.perform = this.perform.bind(this);
  }

  perform() {
    fetch('api:8000/api/command/rainbow', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      }
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Clear success:", response.statusText);
        } else {
          console.log("Clear error");
        }
      },
      (reason) => {
        console.log("Clear error:", reason);
      });
  }

  render() {
    return (
      <div>
        <button style={{padding:10}} onClick={this.perform}>Clear</button>
      </div>
    );
  }
}