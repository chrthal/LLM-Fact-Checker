package llm_fetcher

import (
	"log"

	"chrthal/llm-fact-checker/models"
)

func LLMFetcher(lLLData models.LLMData, query string) models.LLMData {
	switch lLLData.LLM.String() {
	case "GPT4":
		lLLData.Response = gpt4_fetcher(query)
	case "GPT3.5":
		lLLData.Response = gpt3_fetcher(query)
	case "Test-Data":
		lLLData.Response = "The capital of France is Paris."
	default:
		log.Printf("Unknown large language model: %s\n", lLLData.LLM.String())
	}
	return lLLData
}
