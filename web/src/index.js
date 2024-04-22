// index.js

import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import ResolvedJobs from './components/ResolvedJobs';

ReactDOM.render(
  <React.StrictMode>
    <App>
      <ResolvedJobs /> {/* Render the DashboardOverview component within App */}
    </App>
  </React.StrictMode>,
  document.getElementById('root')
);
