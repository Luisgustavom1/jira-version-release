package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/entity"
)

const VERSION_WEBHOOK_NAME = "release-notes-bot"

func SubscribeInWebhook(j *jira.JiraConnect) {
	jiraWebhookSubscribe, err := json.Marshal(map[string]any{
		"name":   VERSION_WEBHOOK_NAME,
		"url":    configs.GetEnv("BOT_URL") + "/webhooks",
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

func ListWebhooks[K entity.JiraWebhook](j *jira.JiraConnect) []K {
	res, err := j.NewJiraRequest(
		"GET",
		"/webhooks/1.0/webhook",
		nil,
	)
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	var versions []K
	err = json.NewDecoder(res.Body).Decode(&versions)
	if err != nil {
		log.Fatalf("Cannot parse response!")
		return nil
	}

	return versions
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
			http.Error(w, "Cannot parse response!", http.StatusBadRequest)
			return
		}

		ch <- jiraWebhookVersion
	})
}
