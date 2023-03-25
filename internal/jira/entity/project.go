package entity

type JiraProject struct {
	Expand         string          `json:"expand"`
	Self           string          `json:"self"`
	ID             string          `json:"id"`
	Key            string          `json:"key"`
	Description    string          `json:"description"`
	IssueTypes     []JiraIssueType `json:"issueTypes"`
	AssigneeType   string          `json:"assigneeType"`
	Versions       []JiraVersion   `json:"versions"`
	Name           string          `json:"name"`
	ProjectTypeKey string          `json:"projectTypeKey"`
	Simplified     bool            `json:"simplified"`
	Style          string          `json:"style"`
	IsPrivate      bool            `json:"isPrivate"`
}
