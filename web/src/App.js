// App.js

import React from 'react';
import DashboardOverview from './components/DashboardOverview';
import { Container } from '@material-ui/core';

const App = () => {
  return (
    <Container maxWidth="md" style={{ marginTop: '20px' }}>
      <h1>Overview</h1>
      <DashboardOverview />
    </Container>
  );
};

export default App;
