package main

import (
	"fmt"
	"net/http"

	"github.com/Luisgustavom1/release-notes-bot/configs"
	"github.com/Luisgustavom1/release-notes-bot/cmd/discord-bot"
)

func init() {
	configs.LoadEnvs()
}

func main() {
	discord_bot.InitDiscordBot()

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
