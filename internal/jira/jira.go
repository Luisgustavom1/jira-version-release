package jira

import (
	"io"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/configs"
)

type JiraConnect struct {
	JiraAuthentication string
}

func Connect(jiraAuthentication string) *JiraConnect {
	return &JiraConnect{
		JiraAuthentication: jiraAuthentication,
	}
}

func (j *JiraConnect) NewJiraRequest(method string, path string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		configs.GetEnv("MY_JIRA_URL")+"/rest"+path,
		body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+j.JiraAuthentication)
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
