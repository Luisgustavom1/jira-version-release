package entity

type JiraWebhook struct {
	Name                   string             `json:"name"`
	URL                    string             `json:"url"`
	ExcludeBody            bool               `json:"excludeBody"`
	Filters                JiraWebhookFilters `json:"filters"`
	Events                 []string           `json:"events"`
	Enabled                bool               `json:"enabled"`
	Self                   string             `json:"self"`
	LastUpdatedDisplayName string             `json:"lastUpdatedDisplayName"`
	LastUpdated            int64              `json:"lastUpdated"`
}

type JiraWebhookFilters struct {
	IssueRelatedEventsSection string `json:"issue-related-events-section"`
}
