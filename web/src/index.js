// index.js

import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import DashboardOverview from './components/DashboardOverview'; // Import the DashboardOverview component

ReactDOM.render(
  <React.StrictMode>
    <App>
      <DashboardOverview /> {/* Render the DashboardOverview component within App */}
    </App>
  </React.StrictMode>,
  document.getElementById('root')
);
