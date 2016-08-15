package ui

// SearchResults is a list of SearchResult
type SearchResults struct {
	Query     SearchQuery `json:"query"`
	TimeTaken int64       `json:"timeTakenInMs,omitempty"`
	Results   interface{} `json:"results,omitempty"`
}

type SearchQuery struct {
	Query             string   `json:"q,omitempty"`
	InTitle           bool     `json:"inTitle,omitempty"`
	Hashtags          []string `json:"hashtags,omitempty"`
	MicrocosmIDsQuery []int64  `json:"forumId,omitempty"`
	ItemTypesQuery    []string `json:"type,omitempty"`
	ItemIDsQuery      []int64  `json:"id,omitempty"`
	ProfileID         int64    `json:"authorId,omitempty"`
	Emails            []string `json:"email,omitempty"`
	Following         bool     `json:"following,omitempty"`
	Since             string   `json:"since,omitempty"`
	Until             string   `json:"until,omitempty"`
	EventAfter        string   `json:"eventAfter,omitempty"`
	EventBefore       string   `json:"eventBefore,omitempty"`
	Attendee          bool     `json:"attendee,omitempty"`
	Has               []string `json:"has,omitempty"`
	Sort              string   `json:"sort,omitempty"`

	Ignored  string `json:"ignored,omitempty"`
	Searched string `json:"searched,omitempty"`
}
