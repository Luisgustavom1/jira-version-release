package discord

import (
	"log"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/bwmarrin/discordgo"
)

func Connect() *discordgo.Session {
	token := configs.GetEnv("DISCORD_BOT_TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Panicln(err)
	}
	
	return discord
}

func SendChannelMessage(discord *discordgo.Session, channel string, message string) {
	channels, err := discord.GuildChannels(configs.GetEnv("DISCORD_GUILD_ID"))
	if err != nil {
		log.Println(err)
		return
	}
	for _, c := range channels {
		if c.Name == channel {
			discord.ChannelMessageSend(c.ID, message)
		}
	}
}
