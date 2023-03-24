package entity

type JiraIssue struct {
	Expand string          `json:"expand"`
	ID     string          `json:"id"`
	Self   string          `json:"self"`
	Key    string          `json:"key"`
	Fields JiraIssueFields `json:"fields"`
}

type JiraIssueFields struct {
	IssueType JiraIssueType `json:"issuetype"`
	Summary   string        `json:"summary"`
}

type JiraIssueType struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}
