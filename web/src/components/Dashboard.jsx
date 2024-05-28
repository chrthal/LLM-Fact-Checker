import React, { useState, useEffect } from "react";
import { addJob, fetchResolvedJobs, fetchStatus } from "../api/jobs";
import JobFormModal from './JobFormModal';
import SettingsModal from './SettingsModal';

export default function Dashboard() {
  const [jobs, setJobs] = useState([]);
  const [config, setConfig] = useState(null);
  const [error, setError] = useState(null);
  const [status, setStatus] = useState({
    queuedJobs: 0,
    runningJobs: 0,
    resolvedJobs: 0,
  });
  const [request, setRequest] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isSettingsModalOpen, setIsSettingsModalOpen] = useState(false); 

  useEffect(() => {
    const getStatus = async () => {
      try {
        const result = await fetchStatus();
        setStatus(result);
      } catch (err) {
        console.error("Axios error:", err.response || err.message);
        setError(err.response || err.message);
      }
    };

    getStatus();
  }, []);

  useEffect(() => {
    const getJobs = async () => {
      try {
        const result = await fetchResolvedJobs();
        setJobs(result);
      } catch (err) {
        console.error("Axios error:", err.response || err.message);
        setError(err.response || err.message);
      }
    };

    getJobs();
  }, []);

  const searchEngineMappings = {
    0: "Google",
    1: "Bing",
    2: "Yahoo",
    3: "DuckDuckGo",
    4: "Baidu",
  };

  const llmMappings = {
    0: "GPT4",
    1: "GPt3.5",
    2: "Gemini",
    3: "Ollama",
  };

  const handleAddJob = async (jobData) => {
    try {
      const newJob = await addJob(jobData);
      // Assuming the API returns the updated job list or the new job
      setRequest({ ...request, jobs: [newJob] }); // Assuming there's a single job per request
      alert("Job added successfully!");
    } catch (error) {
      alert("Failed to add job: " + error.message);
    }
  };
  const handleSettings = async (configData) => {
    try {
      const newConfig = await setConfig(configData);
      setRequest({ ...request, config: [newConfig] }); // Assuming there's a single job per request
      alert("Settings updated successfully!");
    } catch (error) {
      alert("Failed to add job: " + error.message);
    }
  };
  return (
    <div className="min-h-screen bg-zinc-100 dark:bg-zinc-900 p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-semibold text-zinc-900 dark:text-zinc-100">
          Dashboard
        </h1>
        <div className="flex space-x-2">
          <button className="bg-zinc-200 dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 px-4 py-2 rounded" 
            onClick={() => setIsSettingsModalOpen(true)}>
            Settings
          </button>
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded"
            onClick={() => setIsModalOpen(true)}>Add Job
          </button>
        </div>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
        <div className="bg-white dark:bg-zinc-800 p-4 rounded shadow">
          <h2 className="text-lg font-medium text-zinc-900 dark:text-zinc-100">
            Queued Jobs
          </h2>
          <p className="text-2xl font-bold text-zinc-900 dark:text-zinc-100">
            {status.queuedJobs}
          </p>
        </div>
        <div className="bg-white dark:bg-zinc-800 p-4 rounded shadow">
          <h2 className="text-lg font-medium text-zinc-900 dark:text-zinc-100">
            Running
          </h2>
          <p className="text-2xl font-bold text-zinc-900 dark:text-zinc-100">
            {status.runningJobs}
          </p>
        </div>
        <div className="bg-white dark:bg-zinc-800 p-4 rounded shadow">
          <h2 className="text-lg font-medium text-zinc-900 dark:text-zinc-100">
            Resolved Jobs
          </h2>
          <p className="text-2xl font-bold text-zinc-900 dark:text-zinc-100">
            {status.resolvedJobs}
          </p>
        </div>
      </div>
      <div className="bg-white dark:bg-zinc-800 rounded shadow overflow-hidden">
        <table className="min-w-full divide-y divide-zinc-200 dark:divide-zinc-700">
          <thead className="bg-zinc-50 dark:bg-zinc-700">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Number
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Question
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Pages to crawl
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                LLM
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Search Engines
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Similarity
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Start Date
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-zinc-500 dark:text-zinc-300 uppercase tracking-wider">
                Status
              </th>
            </tr>
          </thead>
          <tbody>
            {error && <div>Error: {error}</div>}
            {!jobs ? (
              <div>Loading...</div>
            ) : (
              jobs &&
              jobs.map((job, index) => (
                <tr key={index}>
                  <td className="py-3 px-6 text-white">{job.jobId}</td>
                  <td className="py-3 px-6 text-white">{job.question}</td>
                  <td className="py-3 px-6 text-white">{job.pagesToCrawl}</td>
                  <td className="py-3 px-6 text-white">
                    {job.largeLanguageModelMData
                      .map((llm) => llmMappings[llm.LLM])
                      .join(", ")}
                  </td>
                  <td className="py-3 px-6 text-white">
                    {job.searchEngineData
                      .map(
                        (engine) => searchEngineMappings[engine.SearchEngine]
                      )
                      .join(", ")}
                  </td>
                  <td className="py-3 px-6 text-white">{job.crawledData.similarity}</td>
                  <td className="py-3 px-6 text-white">{job.startDate}</td>
                  <td className="py-3 px-6 text-white">{job.status}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
      <JobFormModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleAddJob}
      />
      <SettingsModal
        isOpen={isSettingsModalOpen}
        onClose={() => setIsSettingsModalOpen(false)}
        onSubmit={handleSettings}
      />
    </div>
  );
}
