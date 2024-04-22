import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/v1';

export const fetchResolvedJobs = async () => {
  return axios.get(`${API_BASE_URL}/resolvedJobs`);
};

export const addJob = async (jobData) => {
  return axios.post(`${API_BASE_URL}/addJob`, jobData);
};
