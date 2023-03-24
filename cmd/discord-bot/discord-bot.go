package discord_bot

import (
	"encoding/base64"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/internal/discord"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/api"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/entity"
)

func InitDiscordBot() {
	jiraApiToken := configs.GetEnv("JIRA_API_TOKEN")
	jiraEmail := configs.GetEnv("JIRA_EMAIL")
	jiraAuthentication := base64.StdEncoding.EncodeToString([]byte(
		jiraEmail + ":" + jiraApiToken,
	))

	jiraConnection := jira.Connect(jiraAuthentication)
	discordSession := discord.Connect()

	if !alreadySubscribeInVersionWebhook(jiraConnection) {
		api.SubscribeInWebhook(jiraConnection)
	}

	jiraVersion := make(chan entity.JiraWebhookVersion)
	go api.ListenWebhook(jiraVersion)

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
}

func alreadySubscribeInVersionWebhook(j *jira.JiraConnect) bool {
	alreadySubscribe := false
	versions := api.ListAllWebhooks(j)

	for _, v := range versions {
		if v.Name == api.VERSION_WEBHOOK_NAME {
			alreadySubscribe = true
		}
	}

	return alreadySubscribe
}
