import React, {Component} from 'react';
import {AlphaPicker, CirclePicker} from 'react-color';
import {Grid} from "semantic-ui-react";
import {default as RadialPicker} from "@radial-color-picker/react-color-picker";
import ColorConvert from "color-convert"

export class ColorPicker extends Component {
  constructor(props) {
    super(props);
    this.state = {
      color: this.props.color
    };
    this.onRadialColorChange = this.onRadialColorChange.bind(this);
    this.onQuickColorChange = this.onQuickColorChange.bind(this);
  }

  onRadialColorChange({ hue, saturation, luminosity, alpha }) {
    let rgb = ColorConvert.hsl.rgb(hue, saturation, luminosity);
    let color = {
      color: {
        rgb: {
          r: rgb[0],
          g: rgb[1],
          b: rgb[2],
          a: alpha
        },
        hsl: {
          h: hue,
          s: 1,
          l: 0.5,
          a: alpha
        }
      }
    };
    this.setState(color);
    this.props.onChange && this.props.onChange(color);
  };

  onQuickColorChange(e) {
    let color = {
      color: {
        rgb: {
          r: e.rgb.r,
          g: e.rgb.g,
          b: e.rgb.b,
          a: e.rgb.a
        },
        hsl: {
          h: e.hsl.h,
          s: e.hsl.s,
          l: e.hsl.l,
          a: e.hsl.a
        }
      }
    };
    this.setState(color);
    this.props.onChange && this.props.onChange(color);
  };

  render() {
    return (
      <div>
        <Grid>
          <Grid.Row centered>
            <RadialPicker
              hue={this.state.color.hsl.h}
              saturation={this.state.color.hsl.s*100}
              luminosity={this.state.color.hsl.l*100}
              alpha={this.state.color.hsl.a}
              onChange={this.onRadialColorChange}/>
          </Grid.Row>
          <Grid.Row centered>
            <CirclePicker
              colors={['#FF0000', '#FF7F00', '#00FF00', '#FFFF00', '#0000FF', '#4B0082', '#8B00FF', '#FFFFFF']}
              onChangeComplete={this.onQuickColorChange}/>
          </Grid.Row>
          <Grid.Row centered>
            <AlphaPicker
              color={{
                h: this.state.color.hsl.h,
                s: this.state.color.hsl.s,
                l: this.state.color.hsl.l,
                a: this.state.color.hsl.a
              }}
              onChange={this.onQuickColorChange}/>
          </Grid.Row>
        </Grid>
      </div>
    );
  }
}