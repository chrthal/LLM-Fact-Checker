// main.js

import React, { useState, useEffect } from "react";
import { List, ListItem, ListItemText, Button } from "@material-ui/core";

const DashboardOverview = () => {
  const [jobs, setJobs] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/resolvedJobs");
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const jsonJobs = await response.json();
        if (!jsonJobs || !Array.isArray(jsonJobs) || jsonJobs.length === 0) {
          throw new Error("Jobs data is empty or not usable");
        }
        setJobs(jsonJobs);
      } catch (error) {
        setError(error.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleAddJob = () => {
    // Add your logic here to handle adding a new job
    console.log("Add job button clicked");
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (jobs.length === 0) {
    return <div>No jobs available</div>;
  }

  return (
    <div>
      <h2>Jobs</h2>
      <table>
        <thead>
          <tr>
            <th>Job ID</th>
            <th>Question</th>
            <th>Pages to Crawl</th>
            <th>Search Engines</th>
            <th>Status</th>
            <th>Start Date</th>
            <th>Last Update</th>
          </tr>
        </thead>
        <tbody>
          {jobs.map((job) => (
            <tr key={job.jobId}>
              <td>{job.jobId}</td>
              <td>{job.question}</td>
              <td>{job.pagesToCrawl}</td>
              <td>{job.searchEngines}</td>
              <td>{job.status}</td>
              <td>{job.startDate}</td>
              <td>{job.lastUpdate}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <Button variant="contained" color="primary" href="/addJob">
        Add New Job
      </Button>
    </div>
  );
};

export default DashboardOverview;
