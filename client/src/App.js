import React from 'react';
import logo from './logo.svg';
import './App.css';

const { HardwareServiceClient } = require('./proto/service_grpc_web_pb');
const { BasicRequest, BasicResponse, FadeRequest, FadeResponse } = require('./proto/service_pb');

var client = new HardwareServiceClient('localhost:50051', null, null);

var callBasic = (color, intensity) => {
  const request = new BasicRequest();
  request.setColor(color);
  request.setIntensity(intensity);
  client.basic(request, {}, (err, response) => {
    if (response == null) {
      console.log(err);
    } else {
      console.log(response)
    }
  });
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <button style={{padding:10}} onClick={callBasic({'r': 255, 'g': 255, 'b':255}, 1)}>Basic</button>
      </header>
    </div>
  );
}

export default App;
