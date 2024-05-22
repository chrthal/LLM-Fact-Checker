import React, { useState } from 'react';
import { setApiKeys } from '../api/jobs'; // Adjust the import path as necessary

const SettingsModal = ({ isOpen, onClose }) => {
  const [formData, setFormData] = useState({
    OPENAI_API_KEY: '',
    GCS_KEY: '',
    GCS_ID: '',
    BING_KEY: '',
    OLLAMA_HOST: '',
    OLLAMA_VERBOSE: '',
  });

  if (!isOpen) {
    return null;
  }

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await setApiKeys(formData);
      console.log('Response:', response); // Log the response
      alert("API keys set successfully!");
      // Resetting the form fields
      setFormData({
        OPENAI_API_KEY: '',
        GCS_KEY: '',
        GCS_ID: '',
        BING_KEY: '',
        OLLAMA_HOST: '',
        OLLAMA_VERBOSE: '',
      });
      onClose();
    } catch (error) {
      console.error('Error setting API keys:', error); // Log the error
      alert("Failed to set API keys: " + error.message);
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white dark:bg-zinc-800 p-4 rounded shadow-md w-full max-w-md">
        <h2 className="text-xl font-semibold text-zinc-900 dark:text-zinc-100 mb-4">Set API Keys</h2>
        <form onSubmit={handleSubmit}>
          {Object.keys(formData).map((key) => (
            <div className="mb-4" key={key}>
              <label className="block text-zinc-900 dark:text-zinc-100 mb-2" htmlFor={key}>
                {key.replace('_', ' ')}
              </label>
              <input
                type="text"
                name={key}
                value={formData[key]}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded dark:bg-zinc-700 dark:text-zinc-100"
                required
              />
            </div>
          ))}
          <div className="flex justify-end space-x-2">
            <button type="button" onClick={onClose} className="bg-gray-200 dark:bg-zinc-600 text-zinc-900 dark:text-zinc-100 px-4 py-2 rounded">
              Cancel
            </button>
            <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">
              Set API Keys
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SettingsModal;
