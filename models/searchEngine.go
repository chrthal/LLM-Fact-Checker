package models

type SearchEngine int

const (
	Google SearchEngine = iota
	Bing
	Yahoo
	DuckDuckGo
	Baidu
)

type SearchEngineData struct {
	SearchEngine SearchEngine
	Urls         []string
}

func (se SearchEngine) String() string {
	return [...]string{"Google", "Bing", "Yahoo", "DuckDuckGo", "Baidu"}[se]
}
