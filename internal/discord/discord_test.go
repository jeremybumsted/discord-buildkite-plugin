package discord

import (
	"bk-plugin/internal/build"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDiscordEmbedFailedJob(t *testing.T) {
	
	// build info fixture used by LoadDiscordEmbed
	mockBuildInfo := build.BuildInfo{
		Number:   "20",
		URL:      "https://my-build-url/",
		Pipeline: "llamas",
		Author:   "carl",
		Branch:   "main",
		Message:  "failed job embed",
	}

	expectedEmbed := []DiscordEmbed{
		{
			Title:       "failed job embed",
			Description: "**[Failed]** llamas (main) #20",
			Color:       16711680,
			URL:         "https://my-build-url/",
		},
	}

	embed := LoadDiscordEmbed("1", mockBuildInfo)

	assert.Equal(t, expectedEmbed, embed)
}

func TestCreateDiscordEmbedPassedJob(t *testing.T) {
	
	mockBuildInfo := build.BuildInfo{
		Number:   "20",
		URL:      "https://my-build-url/",
		Pipeline: "llamas",
		Author:   "carl",
		Branch:   "main",
		Message:  "passed job embed",
	}

	expectedEmbed := []DiscordEmbed{
		{
			Title:       "passed job embed",
			Description: "**[Passed]** llamas (main) #20",
			Color:       65280,
			URL:         "https://my-build-url/",
		},
	}

	embed := LoadDiscordEmbed("0", mockBuildInfo)

	assert.Equal(t, expectedEmbed, embed)
}