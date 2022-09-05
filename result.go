package main

import "strings"

func deleteHighligth(s string) string {
	return strings.Replace(strings.Replace(s, "<gzkNfoUU>", "", -1), "</gzkNfoUU>", "", -1)
}

type AlfredLink struct {
	Link     string
	Subtitle string
	Title    string
	UID      string
}

type SearchResult struct {
	Results              []Result             `json:"results"`
	Total                int64                `json:"total"`
	RecordMap            RecordMap            `json:"recordMap"`
	TrackEventProperties TrackEventProperties `json:"trackEventProperties"`
}

func (r *SearchResult) AlfredLinks(url string) []*AlfredLink {
	links := []*AlfredLink{}
	for _, result := range r.Results {
		subtitle := ""
		if result.Highlight.PathText != nil {
			subtitle = deleteHighligth(*result.Highlight.PathText)
		}
		links = append(links, &AlfredLink{
			UID:      result.ID,
			Title:    deleteHighligth(result.Highlight.Text),
			Subtitle: subtitle,
			Link:     url + strings.Replace(result.ID, "-", "", -1),
		})
	}
	return links
}

type RecordMap struct {
	Version    int64 `json:"__version__"`
	Block      Block `json:"block"`
	Space      Block `json:"space"`
	Collection Block `json:"collection"`
}

type Block struct {
}

type Result struct {
	ID          string    `json:"id"`
	IsNavigable bool      `json:"isNavigable"`
	Highlight   Highlight `json:"highlight"`
	Score       float64   `json:"score"`
	SpaceID     string    `json:"spaceId"`
	Source      Source    `json:"source"`
}

type Highlight struct {
	Text     string  `json:"text"`
	PathText *string `json:"pathText,omitempty"`
}

type TrackEventProperties struct {
	QueryLength               int64             `json:"queryLength"`
	QueryTokensNaive          int64             `json:"queryTokensNaive"`
	TruncatedQueryLength      int64             `json:"truncatedQueryLength"`
	TruncatedQueryTokensNaive int64             `json:"truncatedQueryTokensNaive"`
	NumBlockIDSInQuery        int64             `json:"numBlockIdsInQuery"`
	WorkspaceID               string            `json:"workspaceId"`
	Took                      int64             `json:"took"`
	IndexAlias                string            `json:"indexAlias"`
	Shard                     int64             `json:"shard"`
	Nodes                     []string          `json:"nodes"`
	Language                  string            `json:"language"`
	SearchExperiments         SearchExperiments `json:"searchExperiments"`
	SearchSessionID           string            `json:"searchSessionId"`
	QueryType                 string            `json:"queryType"`
	RequestSource             string            `json:"requestSource"`
}

type SearchExperiments struct {
	SearchNoWildcardForNonCJK string `json:"search-no-wildcard-for-non-cjk"`
	SearchNoFuzziness         string `json:"search-no-fuzziness"`
}

type Source string

const (
	Es Source = "es"
)
