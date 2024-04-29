package llm_fetcher

import (
	"log"
	"regexp"

	"chrthal/llm-fact-checker/models"
)

func LLMFetcher(lLLData models.LLMData, query string) models.LLMData {
	llmResponse := ""
	switch lLLData.LLM.String() {
	case "GPT4":
		llmResponse = gpt4_fetcher(query)
	case "GPT3.5":
		llmResponse = gpt3_fetcher(query)
	case "Ollama":
		llmResponse = ollama_fetcher(query)
	default:
		log.Printf("Unknown large language model: %s\n", lLLData.LLM.String())
	}
	lLLData.Response = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(llmResponse, "")
	return lLLData
}
