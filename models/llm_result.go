package models

type Llm_result struct {
	ID       int    `json:"llmResultId"`
	Response string `json:"response"`
	Engine   string `json:"Engine"`
}
