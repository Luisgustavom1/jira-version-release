package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/discord"
)

type JiraVersionLink struct {
	Self string `json:"self"`
	Name string `json:"name"`
	Link struct {
		GlobalId             string   `json:"globalId"`
		MyCustomLinkProperty bool     `json:"myCustomLinkProperty"`
		Colors               []string `json:"colors"`
	} `json:"link"`
}

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

func init() {
	configs.LoadEnvs()
}

func main() {
	jiraURL := configs.GetEnv("JIRA_URL")
	jiraApiToken := configs.GetEnv("JIRA_API_TOKEN")
	jiraEmail := configs.GetEnv("JIRA_EMAIL")

	jiraAuthentication := base64.StdEncoding.EncodeToString([]byte(
		jiraEmail + ":" + jiraApiToken,
	))

	jiraWebhookSubscribe, err := json.Marshal(map[string]any{
		"name":   "my first webhook via rest",
		"url":    "https://d630-177-106-118-8.sa.ngrok.io/webhooks",
		"events": []string{"jira:version_released"},
	})
	if err != nil {
		return
	}

	discordSession := discord.Connect()

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var jiraWebhookVersion JiraWebhookVersion
		err = json.NewDecoder(r.Body).Decode(&jiraWebhookVersion)
		if err != nil {
			log.Fatalf(err.Error())
			http.Error(w, "Cannot parse response!", http.StatusBadRequest)
			return
		}

		fmt.Println("Webhook sended -> ", jiraWebhookVersion.Version.Description)
		discord.SendChannelMessage(
			discordSession,
			"notes",
			jiraWebhookVersion.Version.Description,
		)
	})

	http.HandleFunc("/register/webhooks/release_notes", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		req, err := http.NewRequest(
			"POST",
			jiraURL+"/rest/webhooks/1.0/webhook",
			bytes.NewBuffer(jiraWebhookSubscribe),
		)
		if err != nil {
			return
		}

		req.Header.Add("Authorization", "Basic "+jiraAuthentication)
		req.Header.Add("Content-Type", "application/json")
		response, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Subscribed -> ", response.StatusCode)
	})

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
