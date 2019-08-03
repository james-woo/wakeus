import React, {Component} from 'react';
import {Button, Input, Segment} from "semantic-ui-react";

export class RainbowForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      disabled: false,
      cycles: 1
    };
    this.handleChangeCycles = this.handleChangeCycles.bind(this);
    this.perform = this.perform.bind(this);
  }

  handleChangeCycles(event) {
    this.setState({
      cycles: parseInt(event.target.value)
    });
  }

  perform() {
    this.setState({disabled: true});
    let host = process.env.REACT_APP_HOST_IP_ADDRESS || "192.168.1.52";
    console.log(`Request fade: 
      cycles: ${this.state.cycles} 
      to host ${host}`
    );
    fetch(`http://${host}:8000/api/command/rainbow`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        cycles: this.state.cycles
      })
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Rainbow success:", response.statusText);
          this.setState({disabled: false});
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
      <Segment className="rainbow">
        <h1>Rainbow</h1>
        <Input
          label='Cycles'
          type='number'
          action={<Button content={'Perform'} onClick={this.perform}/>}
          placeholder='1'
          onChange={this.handleChangeCycles}/>
      </Segment>
    );
  }
}