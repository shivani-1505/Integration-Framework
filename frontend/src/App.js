import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import UserComponent from './components/UserComponent';

function App() {
  return (
    <Router>
      <div>
        <h1>Zapier Clone</h1>
        <Switch>
          <Route path="/users" component={UserComponent} />
          {/* Add more routes as needed */}
        </Switch>
      </div>
    </Router>
  );
}

export default App;