import React, { Component }  from 'react';
import {ChromePicker} from "react-color";

export class FadeForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      color1: '#fff',
      color2: '#fff',
      duration: 5
    };

    this.handleChangeColor1 = this.handleChangeColor1.bind(this);
    this.handleChangeColor2 = this.handleChangeColor2.bind(this);
    this.handleChangeDuration = this.handleChangeDuration.bind(this);
    this.perform = this.perform.bind(this);
  }

  handleChangeColor1(color, event) {
    this.setState({
      color1: color.rgb
    });
    console.log(color.rgb);
  }

  handleChangeColor2(color, event) {
    this.setState({
      color2: color.rgb
    });
    console.log(color.rgb);
  }

  handleChangeDuration(event) {
    this.setState({
      duration: event.target.value
    })
  }

  perform() {
    fetch('http://api:8000/api/command/fade', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        start_color: {
          r: this.state.color1.r,
          g: this.state.color1.g,
          b: this.state.color1.b
        },
        end_color: {
          r: this.state.color1.r,
          g: this.state.color1.g,
          b: this.state.color1.b
        },
        start_intensity: this.state.color1.a,
        end_intensity: this.state.color1.a,
        duration: this.state.duration
      })
    }).then(
      (response) => {
        if (response.ok) {
          console.log("Fade success: " + response.statusText);
        } else {
          console.log("Fade error");
        }
      },
      (reason) => {
        console.log("Fade error: ", reason);
      });
  }

  render() {
    return (
      <div>
        <p>Start Color</p>
        <ChromePicker
          color={this.state.color1}
          onChange={this.handleChangeColor1}
        />
        <p>End Color</p>
        <ChromePicker
          color={this.state.color2}
          onChange={this.handleChangeColor2}
        />
        <p>Duration</p>
        <input type="number" name="duration" value={this.state.duration} onChange={this.handleChangeDuration}/>
        <button style={{padding:10}} onClick={this.perform}>Fade</button>
      </div>
    );
  }
}