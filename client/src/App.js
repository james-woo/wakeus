import React from 'react';
import './App.css';

import {BasicForm} from "./components/BasicForm";
import {ClearForm} from "./components/ClearForm"
import {RainbowForm} from "./components/RainbowForm";
import {FadeForm} from "./components/FadeForm";
import {Link, Route, BrowserRouter as Router} from "react-router-dom";
import {Menu} from "semantic-ui-react";
import { SettingsForm } from './components/SettingsForm';

function App() {
  return (
    <div className="App">
      <Router>
        <Menu stackable inverted>
          <Menu.Item name='basic'>
            <Link to="/basic">Basic</Link>
          </Menu.Item>
          <Menu.Item name='fade'>
            <Link to="/fade">Fade</Link>
          </Menu.Item>
          <Menu.Item name='rainbow'>
            <Link to="/rainbow">Rainbow</Link>
          </Menu.Item>
          <Menu.Item name='settings'>
            <Link to="/settings">Settings</Link>
          </Menu.Item>
        </Menu>
        <Route path="/basic" component={BasicForm} />
        <Route path="/fade" component={FadeForm} />
        <Route path="/rainbow" component={RainbowForm} />
        <Route path="/settings" component={SettingsForm} />
      </Router>
      <ClearForm/>
    </div>
  );
}

export default App;
