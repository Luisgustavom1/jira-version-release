package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/jira"
	"github.com/Luisgustavom1/release-notes-bot/jira/entity"
)

func SubscribeInWebhook(j *jira.JiraConnect) {
	jiraWebhookSubscribe, err := json.Marshal(map[string]any{
		"name":   "release-notes-bot",
		"url":    "https://d630-177-106-118-8.sa.ngrok.io/webhooks",
		"events": []string{"jira:version_released"},
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	res, err := j.NewJiraRequest(
		"POST",
		"/webhooks/1.0/webhook",
		bytes.NewBuffer(jiraWebhookSubscribe),
	)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fmt.Println("Subscribed -> ", res.StatusCode)
}

func ListAllWebhooks(j *jira.JiraConnect) *http.Response {
	res, err := j.NewJiraRequest(
		"GET",
		"/webhooks/1.0/webhook",
		nil,
	)
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	return res
}

func ListenWebhook[K entity.JiraWebhookVersion](ch chan K) {
	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var jiraWebhookVersion K
		err := json.NewDecoder(r.Body).Decode(&jiraWebhookVersion)
		if err != nil {
			log.Fatalf(err.Error())
			http.Error(w, "Cannot parse response!", http.StatusBadRequest)
			return
		}

		ch <- jiraWebhookVersion
	})
}
