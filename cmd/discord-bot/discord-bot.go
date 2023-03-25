package discord_bot

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/internal/discord"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/api"
	"github.com/Luisgustavom1/release-notes-bot/internal/jira/entity"
)

type IssuesGroupByType = map[string][]entity.JiraIssue

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
		for {
			select {
			case webhookVersion := <-jiraVersion:
				issuesSearched, err := api.ListIssuesByVersionId(jiraConnection, webhookVersion.Version.Id)
				if err != nil {
					log.Fatalln(err)
				}

				project, err := api.ListProjectByProjectId(jiraConnection, webhookVersion.Version.ProjectId)
				if err != nil {
					log.Fatalln(err)
				}

				discord.SendChannelMessage(
					discordSession,
					"notes",
					formatReleaseNotesMessage(
						project.Name,
						issuesGroupByType(issuesSearched.Issues),
						webhookVersion.Version,
					),
				)
			}
		}
	}()
}

func alreadySubscribeInVersionWebhook(j *jira.JiraConnect) bool {
	alreadySubscribe := false
	versions := api.ListWebhooks(j)

	for _, v := range versions {
		if v.Name == api.VERSION_WEBHOOK_NAME {
			alreadySubscribe = true
		}
	}

	return alreadySubscribe
}

func formatReleaseNotesMessage(projectName string, issues IssuesGroupByType, version entity.JiraVersion) string {
	header := fmt.Sprintln("# Release notes -", projectName, "-", version.Name)
	description := "\n" + version.Description + "\n"
	issuesList := ""

	for issueType := range issues {
		issuesList += fmt.Sprint("\n", "### ", issueType, "\n\n")

		for _, issue := range issues[issueType] {
			issuesList += fmt.Sprint("[", issue.Key, "]", "(", issue.Fields.IssueType.Self, ") ", issue.Fields.Summary, "\n\n")
		}
	}

	return fmt.Sprintln(header, description, issuesList)
}

func issuesGroupByType[T []entity.JiraIssue](issues T) IssuesGroupByType {
	groupedIssues := IssuesGroupByType{}

	for _, issue := range issues {
		groupedIssues[issue.Fields.IssueType.Name] = append(groupedIssues[issue.Fields.IssueType.Name], issue)
	}

	return groupedIssues
}
