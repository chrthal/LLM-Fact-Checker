package models

type Config struct {
	OpenAIKey     string `json:"OPENAI_API_KEY"`
	GCSKey        string `json:"GCS_KEY"`
	GCSID         string `json:"GCS_ID"`
	BingKey       string `json:"BING_KEY"`
	OllamaHost    string `json:"OLLAMA_HOST"`
	OllamaVerbose string `json:"OLLAMA_VERBOSE"`
}
