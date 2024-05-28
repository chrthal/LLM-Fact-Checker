import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const fetchStatus = async () => {
  try {
    const response = await apiClient.get('/status'); 
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchResolvedJobs = async () => {
  try {
    const response = await apiClient.get('/resolvedJobs');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const addJob = async (jobData) => {
  try {
    console.log('Request payload:', jobData); 
    const response = await apiClient.post('/addJob', jobData); 
    console.log('Response data:', response.data); 
    return response.data;
  } catch (error) {
    console.error('Error adding job:', error); // Log the error
    throw error;
  }
}

export const setApiKeys = async (apiKeys) => {
  try {
    console.log('Set config', apiKeys);
    const response = await apiClient.post('/config', apiKeys); 
    return response.data;
  } catch (error) {
    console.error('Error setting API keys:', error); // Log the error
    throw error;
  }
};
