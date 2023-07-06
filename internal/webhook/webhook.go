package webhook

import (
	"os"
	"encoding/json"
	"bytes"
	"net/http"
	"errors"

	"github.com/charmbracelet/log"
	"bk-plugin/internal/discord"
)

type WebhookPayload struct {
	Content	string `json:"content,omitempty"`
	Embeds	[]discord.DiscordEmbed `json:"embeds,omitempty"`
}

type ConfiguredWebhook struct {
	Payload WebhookPayload
	Url string
}



// Configure the webhook URL from the plugin config
func LoadWebhookConfig() (string, error) {

	var w string

	webhookEnv, exists := os.LookupEnv("BUILDKITE_PLUGIN_DISCORD_NOTIFICATION_WEBHOOK_ENV")
	if exists {
		w, exists = os.LookupEnv(webhookEnv)
		if exists {
			return w, nil
		}
	} else {
			w, exists = os.LookupEnv("DISCORD_WEBHOOK_URL")
		if !exists {
			log.Error("🚨 No Webhook URL set")
			return "", errors.New("🚨 No Webhook URL set")
		}
	}

	return w, nil
}

// Method for sending the payload to the webhook url
func SendWebhook(client *http.Client, hook ConfiguredWebhook) error {
	
	payloadBytes, err := json.Marshal(hook.Payload)
	if err != nil {
		log.Error(err)
		return err
	}

	
	req, err := http.NewRequest("POST", hook.Url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Error(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}


// Process the webhook payload to be sent to Discord, then send it
func ProcessWebhook(embed []discord.DiscordEmbed) error {

	// this is the customizable [message]in the plugin configuration
	content := os.Getenv("BUILDKITE_PLUGIN_DISCORD_NOTIFICATION_MESSAGE")
	
	// Make sure we have a URL to send the hook to
	webhookUrl, err := LoadWebhookConfig()
	if err != nil {
		log.Errorf("Failed to get Webhook URL: %v", err)
		os.Exit(1)
	}
	
	payload := WebhookPayload{
		Content: content,
		Embeds: embed,
	}

	hook := ConfiguredWebhook{
		Payload: payload,
		Url: webhookUrl,
	}

	// create an http client
	client := new(http.Client)

	err = SendWebhook(client, hook)
	if err != nil {
		log.Errorf("Unable to send webhook: %v", err)
		os.Exit(1)
	}

	return nil
}