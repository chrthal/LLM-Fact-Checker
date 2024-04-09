package models

type LLM int

const (
	GPT4 LLM = iota
	GPT3
	Gemini
	TestData
)

func (llm LLM) String() string {
	return [...]string{"GPT4", "GPT3.5", "Gemini", "TestData"}[llm]
}
