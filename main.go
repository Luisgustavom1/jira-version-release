package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/discord"
	"github.com/Luisgustavom1/release-notes-bot/jira"
	"github.com/Luisgustavom1/release-notes-bot/jira/entity"
)

func init() {
	configs.LoadEnvs()
}

func main() {
	jiraApiToken := configs.GetEnv("JIRA_API_TOKEN")
	jiraEmail := configs.GetEnv("JIRA_EMAIL")
	jiraAuthentication := base64.StdEncoding.EncodeToString([]byte(
		jiraEmail + ":" + jiraApiToken,
	))

	jiraConnection := jira.Connect(jiraAuthentication)
	discordSession := discord.Connect()

	http.HandleFunc("/register/webhooks/release_notes", func(w http.ResponseWriter, r *http.Request) {
		jira.SubscribeInWebhook(jiraConnection)
	})

	jiraVersion := make(chan entity.JiraWebhookVersion)
	go jira.ListenWebhook(jiraVersion)

	go func() {
		select {
		case version := <-jiraVersion:
			discord.SendChannelMessage(
				discordSession,
				"notes",
				version.Version.Description,
			)
		}
	}()

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
