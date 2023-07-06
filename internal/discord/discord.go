package discord

import (
	"fmt"
	"bk-plugin/internal/build"
)

type DiscordEmbed struct {
	Title       string		          `json:"title,omitempty"`
	Description string                `json:"description,omitempty"`
	URL         string                `json:"url,omitempty"`
	Color       int                   `json:"color,omitempty"`
	Author      DiscordEmbedAuthor    `json:"author,omitempty"`
	// Fields      []DiscordEmbedField   `json:"fields,omitempty"`
}
type DiscordEmbedAuthor struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}


// Creates the discord embed to be sent in the payload
func LoadDiscordEmbed(status string, b build.BuildInfo) []DiscordEmbed {
	color := 65280
	commandState := "Passed"
	
	if status != "0" {
		commandState = "Failed"
		color = 16711680
	}
	e := []DiscordEmbed{
		{
			Title: b.Message,
			Description: fmt.Sprintf("**[%v]** %v (%v) #%v", commandState, b.Pipeline, b.Branch, b.Number ),
			Color: color,
			URL: b.URL,
		},
	}

	return e
}

