import React, {Component} from 'react';
import {Button, Grid, Input, Segment, Select,} from "semantic-ui-react";
import {ColorPicker} from "./ColorPicker";

export class FadeForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      disabled: false,
      color1: {
        rgb: {r: 255, g: 0, b: 4, a: 1},
        hsl: {h: 360, s: 1, l: 0.5, a: 1}
      },
      color2: {
        rgb: {r: 4, g: 0, b: 255, a: 1},
        hsl: {h: 241, s: 1, l: 0.5, a: 1}
      },
      duration: 5,
      units: 'seconds'
    };

    this.handleChangeColor1 = this.handleChangeColor1.bind(this);
    this.handleChangeColor2 = this.handleChangeColor2.bind(this);
    this.handleChangeDuration = this.handleChangeDuration.bind(this);
    this.handleChangeUnits = this.handleChangeUnits.bind(this);
    this.perform = this.perform.bind(this);
  }

  handleChangeColor1(e) {
    this.setState({
      color1: e.color
    });
  }

  handleChangeColor2(e) {
    this.setState({
      color2: e.color
    });
  }

  handleChangeDuration(event) {
    this.setState({
      duration: event.target.value
    });
  }

  handleChangeUnits(event) {
    this.setState({
      units: event.target.value
    });
  }

  perform() {
    this.setState({disabled: true});
    let host = process.env.REACT_APP_HOST_IP_ADDRESS || "192.168.1.52";
    console.log(`Request fade: 
      start_color: ${JSON.stringify(this.state.color1.rgb)}, 
      end_color: ${JSON.stringify(this.state.color2.rgb)}, 
      start_intensity: ${this.state.color1.rgb.a}, 
      end_intensity: ${this.state.color2.rgb.a}, 
      duration: ${this.state.duration * (this.state.units === 'seconds' ? 1000 : 60000)} 
      to host ${host}`
    );
    fetch(`http://${host}:8000/api/command/fade`, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        start_color: {
          r: this.state.color1.rgb.r,
          g: this.state.color1.rgb.g,
          b: this.state.color1.rgb.b
        },
        end_color: {
          r: this.state.color2.rgb.r,
          g: this.state.color2.rgb.g,
          b: this.state.color2.rgb.b
        },
        start_intensity: this.state.color1.rgb.a,
        end_intensity: this.state.color2.rgb.a,
        duration: this.state.duration * (this.state.units === 'seconds' ? 1000 : 60000)
      })
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Fade success: " + response.statusText);
          this.setState({disabled: false});
        } else {
          console.log("Fade error");
        }
      },
      (reason) => {
        console.log("Fade error: ", reason);
      });
  }

  render() {
    const options = [
      {key: 'seconds', text: 'Seconds', value: 'seconds'},
      {key: 'minutes', text: 'Minutes', value: 'minutes'},
    ];
    return (
      <Segment inverted>
        <h1>Fade</h1>
        <Grid stackable columns={2}>
          <Grid.Row>
            <Grid.Column>
              <p>Start Color</p>
              <ColorPicker
                color={this.state.color1}
                onChange={this.handleChangeColor1}/>
            </Grid.Column>
            <Grid.Column>
              <p>End Color</p>
              <ColorPicker
                color={this.state.color2}
                onChange={this.handleChangeColor2}/>
            </Grid.Column>
          </Grid.Row>
          <Grid.Row columns={1}>
            <Grid.Column>
              <Input type='number' placeholder='5' onChange={this.handleChangeDuration} action>
                <input/>
                <Select compact options={options} defaultValue='seconds' onChange={this.handleChangeUnits}/>
                <Button content='Perform' onClick={this.perform}/>
              </Input>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Segment>
    );
  }
}