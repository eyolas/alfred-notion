package main

type Query struct {
	Type    string  `json:"type"`
	Query   string  `json:"query"`
	SpaceID string  `json:"spaceId"`
	Limit   int64   `json:"limit"`
	Filters Filters `json:"filters"`
	Sort    string  `json:"sort"`
	Source  string  `json:"source"`
}

type Filters struct {
	IsDeletedOnly             bool          `json:"isDeletedOnly"`
	ExcludeTemplates          bool          `json:"excludeTemplates"`
	IsNavigableOnly           bool          `json:"isNavigableOnly"`
	RequireEditPermissions    bool          `json:"requireEditPermissions"`
	Ancestors                 []interface{} `json:"ancestors"`
	CreatedBy                 []interface{} `json:"createdBy"`
	EditedBy                  []interface{} `json:"editedBy"`
	LastEditedTime            interface{}   `json:"lastEditedTime"`
	CreatedTime               interface{}   `json:"createdTime"`
	NavigableBlockContentOnly bool          `json:"navigableBlockContentOnly"`
	InTeams                   []interface{} `json:"inTeams"`
}

func NewSearchQuery(search string, spaceID string) Query {
	return Query{
		Type:    "BlocksInSpace",
		Query:   search,
		SpaceID: spaceID,
		Limit:   9,
		Filters: Filters{
			IsDeletedOnly:             false,
			ExcludeTemplates:          true,
			IsNavigableOnly:           true,
			RequireEditPermissions:    false,
			NavigableBlockContentOnly: true,
			Ancestors:                 []interface{}{},
			CreatedBy:                 []interface{}{},
			EditedBy:                  []interface{}{},
			LastEditedTime:            map[string]interface{}{},
			CreatedTime:               map[string]interface{}{},
			InTeams:                   []interface{}{},
		},
		Sort:   "Relevance",
		Source: "quick_find_input_change",
	}
}
