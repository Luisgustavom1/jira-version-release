package api

import (
	"encoding/json"

	"github.com/Luisgustavom1/release-notes-bot/internal/jira"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/entity"
)

type SearchPaginated struct {
	Expand     string             `json:"expand"`
	StartAt    int                `json:"startAt"`
	MaxResults int                `json:"maxResults"`
	Total      int                `json:"total"`
	Issues     []entity.JiraIssue `json:"issues"`
}

func ListIssuesByVersionId(j *jira.JiraConnect, versionId string) (SearchPaginated, error) {
	res, err := j.NewJiraRequest(
		"GET",
		"/api/2/search?jql=fixVersion="+versionId,
		nil,
	)
	if err != nil {
		return SearchPaginated{}, err
	}

	var issues SearchPaginated
	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return SearchPaginated{}, err
	}

	return issues, nil
}
