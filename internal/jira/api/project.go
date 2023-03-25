package api

import (
	"encoding/json"
	"strconv"

	"github.com/Luisgustavom1/release-notes-bot/internal/jira"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/entity"
)

func ListProjectByProjectId(j *jira.JiraConnect, projectId int) (entity.JiraProject, error) {
	res, err := j.NewJiraRequest(
		"GET",
		"/api/2/project/"+strconv.Itoa(projectId),
		nil,
	)
	if err != nil {
		return entity.JiraProject{}, err
	}

	var project entity.JiraProject
	err = json.NewDecoder(res.Body).Decode(&project)
	if err != nil {
		return entity.JiraProject{}, err
	}

	return project, nil
}
