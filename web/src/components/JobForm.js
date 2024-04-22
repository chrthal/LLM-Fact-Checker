import React, { useState } from "react";
import {
  Button,
  TextField,
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  OutlinedInput,
  Chip,
} from "@material-ui/core";
import { addJob } from "../api/jobApi";

export default function JobForm() {
  const [question, setQuestion] = useState("");
  const [pagesToCrawl, setPagesToCrawl] = useState("");
  const [searchEngines, setSearchEngines] = useState([]);
  const [llm, setLlm] = useState("");

  const searchEngineOptions = [
    "Google",
    "Bing",
    "Yahoo",
    "DuckDuckGo",
    "Baidu",
  ];
  const llmOptions = ["GPT4", "GPT3.5", "Gemini", "TestData"];

  const llmMappings = {
    GPT4: 0,
    "GPT3.5": 1,
    Gemini: 2,
    TestData: 3,
  };

  const searchEngineMappings = {
    Google: 0,
    Bing: 1,
    Yahoo: 2,
    DuckDuckGo: 3,
    Baidu: 4,
  };

  const handleSearchEngineChange = (event) => {
    setSearchEngines(event.target.value);
  };

  const handleLlmChange = (event) => {
    setLlm(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const jobData = {
      question,
      pagesToCrawl: parseInt(pagesToCrawl, 10),
      searchEngines: searchEngines.map(
        (engine) => searchEngineMappings[engine]
      ),
      largeLanguageModels: [llmMappings[llm]],
    };
    // Assuming addJob function sends data to your API
    try {
      await addJob(jobData);
      alert("Job added successfully!");
      // Resetting the form fields
      setQuestion("");
      setPagesToCrawl("");
      setSearchEngines([]);
      setLlm("");
    } catch (error) {
      alert("Failed to add job: " + error.message);
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
      <TextField
        margin="normal"
        required
        fullWidth
        label="Question"
        value={question}
        onChange={(e) => setQuestion(e.target.value)}
      />
      <TextField
        margin="normal"
        required
        fullWidth
        label="Pages to Crawl"
        type="number"
        value={pagesToCrawl}
        onChange={(e) => setPagesToCrawl(e.target.value)}
      />

      <FormControl fullWidth sx={{ mt: 2 }}>
        <InputLabel id="search-engine-select-label">Search Engines</InputLabel>
        <Select
          labelId="search-engine-select-label"
          id="search-engine-select"
          multiple
          value={searchEngines}
          onChange={handleSearchEngineChange}
          input={
            <OutlinedInput id="select-multiple-chip" label="Search Engines" />
          }
          renderValue={(selected) => (
            <Box sx={{ display: "flex", flexWrap: "wrap", gap: 0.5 }}>
              {selected.map((value) => (
                <Chip key={value} label={value} />
              ))}
            </Box>
          )}
        >
          {searchEngineOptions.map((name) => (
            <MenuItem key={name} value={name}>
              {name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <FormControl fullWidth sx={{ mt: 2 }}>
        <InputLabel id="llm-select-label">LLM</InputLabel>
        <Select
          labelId="llm-select-label"
          id="llm-select"
          value={llm}
          label="LLM"
          onChange={handleLlmChange}
        >
          {llmOptions.map((option) => (
            <MenuItem key={option} value={option}>
              {option}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
        Add Job
      </Button>
    </Box>
  );
}
