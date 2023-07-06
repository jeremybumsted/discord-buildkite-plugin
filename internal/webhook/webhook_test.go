package webhook

import (
	"bk-plugin/internal/discord"
	"os"
	"testing"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/h2non/gock"
)

func TestLoadWebhookConfigPluginEnv(t *testing.T) {
	os.Setenv("BUILDKITE_PLUGIN_DISCORD_NOTIFICATION_WEBHOOK_ENV", "LLAMAS")
	os.Setenv("LLAMAS", "https://my-webhook-url")

	expectedUrl := "https://my-webhook-url"

	webhookUrl, err := LoadWebhookConfig()

	assert.Nil(t, err)
	assert.Equal(t, expectedUrl, webhookUrl)

	os.Unsetenv("BUILDKITE_PLUGIN_DISCORD_NOTIFICATION_WEBHOOK_ENV")
	os.Unsetenv("LLAMAS")
}

func TestLoadWebhookConfigDefaultURLEnv(t *testing.T) {
	os.Setenv("DISCORD_WEBHOOK_URL", "https://alpaca-industries/")

	expectedUrl := "https://alpaca-industries/"

	webhookUrl, err := LoadWebhookConfig()

	assert.Nil(t, err)
	assert.Equal(t, expectedUrl, webhookUrl)
	os.Unsetenv("DISCORD_WEBHOOK_URL")
}

func TestLoadWebhookConfigNoUrlSet(t *testing.T) {
	_, err := LoadWebhookConfig()

	assert.Error(t, err, "🚨 No Webhook URL set")
}


func TestSendWebhook(t *testing.T) {

	hook := ConfiguredWebhook{
		Payload: WebhookPayload{
				Content: "From Buildkite with Love",
				Embeds: []discord.DiscordEmbed{
					{
						Title:       "passed job embed",
						Description: "**[Passed]** llamas (main) #20",
						Color:       65280,
						URL:         "https://my-build-url/builds/20",
					},
				},
		},
		Url: "http://discord.com/webhooks",
	}

	defer gock.Off()

	gock.New("http://discord.com").
		Post("webhooks").
		Reply(204)

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	result := SendWebhook(client, hook)

	assert.NoError(t, result)
}