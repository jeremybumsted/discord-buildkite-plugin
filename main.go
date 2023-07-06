package main

import (
	"os"

	"github.com/charmbracelet/log"

	"bk-plugin/internal/build"
	"bk-plugin/internal/discord"
	"bk-plugin/internal/webhook"
)




func main() {
	var embed []discord.DiscordEmbed
	var buildData build.BuildInfo

	jobStatus := os.Getenv("BUILDKITE_COMMAND_EXIT_STATUS")
	buildData = build.GetBuildInfo()

	// Create the Discord embed to be sent
	embed = discord.LoadDiscordEmbed(jobStatus, buildData)
	
	// Do the thing
	err := webhook.ProcessWebhook(embed)
	 if err != nil {
		log.Fatalf("Failed to send webhook event: %s", err)
	}
	log.Info("Sent Webhook Successfully 🎉")
}
