import React, { useState, useEffect } from 'react';
import { fetchResolvedJobs } from '../api/jobApi';
import { List, ListItem, ListItemText, Typography } from '@material-ui/core';

export default function ResolvedJobs() {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const getJobs = async () => {
      const response = await fetchResolvedJobs();
      setJobs(response.data);
    };
    getJobs();
  }, []);

  // Search Engine and LLM mappings as per your backend setup
  const searchEngineMappings = {
    0: 'Google',
    1: 'Bing',
    2: 'Yahoo',
    3: 'DuckDuckGo',
    4: 'Baidu'
  };

  const llmMappings = {
    0: 'GPT4',
    1: 'GPt3.5',
    2: 'Gemini',
    3: 'TestData'
  };

  return (
    <List>
      {jobs.map((job) => (
        <ListItem key={job.jobId} divider>
          <ListItemText
            primary={`Job ID: ${job.jobId} - ${job.question}`}
            secondary={`Pages to Crawl: ${job.pagesToCrawl}, Search Engines Used: ${job.searchEngineData.map(engine => searchEngineMappings[engine.SearchEngine]).join(', ')}`}
          />
          <Typography component="p" variant="body2">
            LLM Used: {job.largeLanguageModelMData.map(llm => llmMappings[llm.LLM]).join(', ')}
          </Typography>
        
          <Typography component="p" variant="caption">
            Status: {job.status}, Last Updated: {new Date(job.lastUpdate).toLocaleString()}
          </Typography>
        </ListItem>
      ))}
    </List>
  );
}
