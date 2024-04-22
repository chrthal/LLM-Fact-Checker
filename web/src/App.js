import React, { useState } from 'react';
import JobForm from './components/JobForm';
import ResolvedJobs from './components/ResolvedJobs';
import { Container, Typography, Box, Button } from '@material-ui/core';

function App() {
  const [showForm, setShowForm] = useState(false);

  const handleToggle = () => {
    setShowForm(!showForm);
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" gutterBottom>
          Job Management
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={handleToggle}
          sx={{ position: 'fixed', bottom: 20, right: 20 }}
        >
          {showForm ? 'View Resolved Jobs' : 'Add New Job'}
        </Button>
        {showForm ? <JobForm /> : <ResolvedJobs />}
      </Box>
    </Container>
  );
}

export default App;
