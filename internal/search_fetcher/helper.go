package search_fetcher

import (
	net_url "net/url"
	"strings"
)

func valid(url string) bool {
	_, err := net_url.ParseRequestURI(url)
	if err != nil || containsYoutube(url) {
		return false
	}
	u, err := net_url.Parse(url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func containsYoutube(url string) bool {
	return strings.Contains(url, "youtube")
}
