import React, { useState } from 'react';
import { addJob } from '../api/jobs';

const JobFormModal = ({ isOpen, onClose }) => {
  const [formData, setFormData] = useState({
    question: '',
    pagesToCrawl: '',
    largeLanguageModel: '', // Single selection for LLM
    searchEngines: [],
  });

  const llmOptions = {
    GPT4: 0,
    "GPT3.5": 1,
    Gemini: 2,
    Ollama: 3,
  };

  const searchEngineOptions = {
    Google: 0,
    Bing: 1,
  };

  if (!isOpen) {
    return null;
  }

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name === 'searchEngines') {
      const options = Array.from(e.target.selectedOptions, option => option.value);
      setFormData({
        ...formData,
        [name]: options,
      });
    } else {
      setFormData({
        ...formData,
        [name]: value,
      });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const jobData = {
      question: formData.question,
      pagesToCrawl: parseInt(formData.pagesToCrawl, 10),
      searchEngines: formData.searchEngines.map(engine => parseInt(engine, 10)),
      largeLanguageModels: [parseInt(formData.largeLanguageModel, 10)], // Single LLM as an array with one item
    };

    console.log('Submitting job data:', jobData); // Log the job data

    try {
      const response = await addJob(jobData);
      console.log('Response:', response); // Log the response
      alert("Job added successfully!");
      // Resetting the form fields
      setFormData({
        question: '',
        pagesToCrawl: '',
        largeLanguageModel: '',
        searchEngines: [],
      });
      onClose();
    } catch (error) {
      console.error('Error adding job:', error); // Log the error
      alert("Failed to add job: " + error.message);
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white dark:bg-zinc-800 p-4 rounded shadow-md w-full max-w-md">
        <h2 className="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4">Add New Job</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-zinc-900 dark:text-zinc-100 mb-2" htmlFor="question">Question</label>
            <input
              type="text"
              name="question"
              value={formData.question}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded dark:bg-zinc-700 dark:text-zinc-100"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-zinc-900 dark:text-zinc-100 mb-2" htmlFor="pagesToCrawl">Pages to Crawl</label>
            <input
              type="number"
              name="pagesToCrawl"
              value={formData.pagesToCrawl}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded dark:bg-zinc-700 dark:text-zinc-100"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-zinc-900 dark:text-zinc-100 mb-2" htmlFor="largeLanguageModel">Large Language Model</label>
            <select
              name="largeLanguageModel"
              value={formData.largeLanguageModel}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded dark:bg-zinc-700 dark:text-zinc-100"
              required
            >
              <option value="">Select LLM</option>
              {Object.keys(llmOptions).map((key) => (
                <option key={key} value={llmOptions[key]}>{key}</option>
              ))}
            </select>
          </div>
          <div className="mb-4">
            <label className="block text-zinc-900 dark:text-zinc-100 mb-2" htmlFor="searchEngines">Search Engines</label>
            <select
              name="searchEngines"
              value={formData.searchEngines}
              onChange={handleChange}
              multiple
              className="w-full px-3 py-2 border rounded dark:bg-zinc-700 dark:text-zinc-100"
              required
            >
              {Object.keys(searchEngineOptions).map((key) => (
                <option key={key} value={searchEngineOptions[key]}>{key}</option>
              ))}
            </select>
          </div>
          <div className="flex justify-end space-x-2">
            <button type="button" onClick={onClose} className="bg-gray-200 dark:bg-zinc-600 text-zinc-900 dark:text-zinc-100 px-4 py-2 rounded">Cancel</button>
            <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">Add Job</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default JobFormModal;
