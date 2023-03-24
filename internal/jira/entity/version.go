package entity

type JiraVersion struct {
	Self            string `json:"self"`
	Id              string `json:"id"`
	Description     string `json:"description"`
	Name            string `json:"name"`
	Archived        bool   `json:"archived"`
	Released        bool   `json:"released"`
	ReleaseDate     string `json:"releaseDate"`
	Overdue         bool   `json:"overdue"`
	UserReleaseDate string `json:"userReleaseDate"`
	ProjectId       int    `json:"projectId"`
}

type JiraWebhookVersion struct {
	Timestamp    float64     `json:"timestamp"`
	WebhookEvent string      `json:"webhookEvent"`
	Version      JiraVersion `json:"version"`
}