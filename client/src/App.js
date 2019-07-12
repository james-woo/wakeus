import React from 'react';
import './App.css';

let basic = () => {
  let request = fetch('http://localhost:8000/api/command/basic', {
    method: 'POST',
    cors: 'enabled',
    headers: {
      'Access-Control-Request-Headers': 'accepts, content-type',
      'Access-Control-Request-Method': 'POST',
      'Access-Control-Allow-Origin': '*',
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      color: {
        r: 100,
        g: 100,
        b: 100
      },
      intensity: 1,
    })
  });
  request.then(
    (response) => {
      if (response.ok) {
        console.log("Basic success: " + response);
      } else {
        console.log("Basic error");
      }
  },
    (reason) => {
      console.log("Basic error: ", reason);
  });
};

let fade = () => {
  fetch('http://localhost:8000/api/command/fade', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      start_color: {
        r: 100,
        g: 100,
        b: 100
      },
      end_color: {
        r: 200,
        g: 100,
        b: 100
      },
      start_intensity: 0,
      end_intensity: 1,
      duration: 6000
    })
  })
};

let clear = () => {
  fetch('http://localhost:8000/api/command/clear', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  })
};

function App() {

  return (
    <div className="App">
      <header className="App-header">
        <button style={{padding:10}} onClick={basic}>Basic</button>
        <button style={{padding:10}} onClick={fade}>Fade</button>
        <button style={{padding:10}} onClick={clear}>Clear</button>
      </header>
    </div>
  );
}

export default App;
