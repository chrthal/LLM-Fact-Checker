import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const fetchStatus = async () => {
  try {
    const response = await apiClient.get('/status'); // Adjust the endpoint as necessary
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchResolvedJobs = async () => {
  try {
    const response = await apiClient.get('/resolvedJobs'); // Adjust the endpoint as necessary
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const addJob = async (jobData) => {
  try {
    console.log('Request payload:', jobData); // Log the request payload
    const response = await apiClient.post('/addJob', jobData); // Adjust the endpoint as necessary
    console.log('Response data:', response.data); // Log the response data
    return response.data;
  } catch (error) {
    console.error('Error adding job:', error); // Log the error
    throw error;
  }
}
