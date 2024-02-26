package search_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type BingAnswer struct {
	Type         string `json:"_type"`
	QueryContext struct {
		OriginalQuery string `json:"originalQuery"`
	} `json:"queryContext"`
	WebPages struct {
		WebSearchURL          string `json:"webSearchUrl"`
		TotalEstimatedMatches int    `json:"totalEstimatedMatches"`
		Value                 []struct {
			ID               string    `json:"id"`
			Name             string    `json:"name"`
			URL              string    `json:"url"`
			IsFamilyFriendly bool      `json:"isFamilyFriendly"`
			DisplayURL       string    `json:"displayUrl"`
			Snippet          string    `json:"snippet"`
			DateLastCrawled  time.Time `json:"dateLastCrawled"`
			SearchTags       []struct {
				Name    string `json:"name"`
				Content string `json:"content"`
			} `json:"searchTags,omitempty"`
			About []struct {
				Name string `json:"name"`
			} `json:"about,omitempty"`
		} `json:"value"`
	} `json:"webPages"`
	RelatedSearches struct {
		ID    string `json:"id"`
		Value []struct {
			Text         string `json:"text"`
			DisplayText  string `json:"displayText"`
			WebSearchURL string `json:"webSearchUrl"`
		} `json:"value"`
	} `json:"relatedSearches"`
	RankingResponse struct {
		Mainline struct {
			Items []struct {
				AnswerType  string `json:"answerType"`
				ResultIndex int    `json:"resultIndex"`
				Value       struct {
					ID string `json:"id"`
				} `json:"value"`
			} `json:"items"`
		} `json:"mainline"`
		Sidebar struct {
			Items []struct {
				AnswerType string `json:"answerType"`
				Value      struct {
					ID string `json:"id"`
				} `json:"value"`
			} `json:"items"`
		} `json:"sidebar"`
	} `json:"rankingResponse"`
}

func BingSearch(query string, numResults int) []string {
	var urls []string

	const endpoint = "https://api.bing.microsoft.com/v7.0/search"
	token := os.Getenv("BING_KEY")

	println("Bing search for:", query)

	// Declare a new GET request.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	// Add the payload to the request.
	param := req.URL.Query()
	param.Add("q", query)
	param.Add("responseFilter", "Webpages")
	param.Add("count", fmt.Sprint(numResults))

	req.URL.RawQuery = param.Encode()

	// Insert the request header.
	req.Header.Add("Ocp-Apim-Subscription-Key", token)

	// Create a new client.
	client := new(http.Client)

	// Send the request to Bing.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	println("Bing search status:", resp.Body)

	// Close the response.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Create a new answer.
	ans := new(BingAnswer)
	err = json.Unmarshal(body, &ans)
	if err != nil {
		fmt.Println(err)
	}

	// Iterate over search results and print the result name and URL.
	for result_index, result := range ans.WebPages.Value {
		fmt.Println(result_index, result.Name, "||", result.URL)
		urls = append(urls, result.URL)
	}

	return urls

}
