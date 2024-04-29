package llm_fetcher

import (
	"fmt"

	"github.com/xyproto/ollamaclient/v2"
)

func ollama_fetcher(question string) string {
	oc := ollamaclient.New()

	if err := oc.PullIfNeeded(); err != nil {
		fmt.Println("Error:", err)
		return err.Error()
	}

	output, err := oc.GetOutput(question)
	if err != nil {
		fmt.Println("Error:", err)
		return err.Error()
	}
	fmt.Printf("\n%s\n", output)

	return output
}
