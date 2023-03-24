package jira

import (
	"io"
	"net/http"
)

type JiraConnect struct {
	JiraAuthentication string
}

func Connect(jiraAuthentication string) *JiraConnect {
	return &JiraConnect{
		JiraAuthentication: jiraAuthentication,
	}
}

func (j *JiraConnect) NewJiraRequest(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		url,
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
