import React from 'react';
import './App.css';

import {BasicForm} from "./components/BasicForm";
import {ClearForm} from "./components/ClearForm"
import {RainbowForm} from "./components/RainbowForm";
import {FadeForm} from "./components/FadeForm";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <BasicForm/>
        <FadeForm/>
        <RainbowForm/>
        <ClearForm/>
      </header>
    </div>
  );
}

export default App;
