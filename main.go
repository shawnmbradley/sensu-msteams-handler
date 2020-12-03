package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// shawn

// Config represents the handler plugin config.
type Config struct {
	sensu.PluginConfig
	Example    string
	WebHookURL string
}

const (
	webHookURL = "webHookURL"
)

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-msteams-handler",
			Short:    "Send alerts to Microsoft Teams channels using webhooks",
			Keyspace: "sensu.io/plugins/sensu-msteams-handler/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "example",
			Env:       "HANDLER_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example string configuration option",
			Value:     &plugin.Example,
		},
		{
			Path:      webHookURL,
			Env:       "WEBHOOK_URL",
			Argument:  "webHookURL",
			Shorthand: "u",
			Default:   "",
			Usage:     "The web hook URL generated by Microsoft Teams.",
			Value:     &plugin.WebHookURL,
		},
	}
)

func main() {
	handler := sensu.NewGoHandler(&plugin.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	log.Println("CHECKING CONFIG")
	if len(plugin.WebHookURL) == 0 {
		return fmt.Errorf("--webHookURL or WEBHOOK_URL environment variable is required")
	}
	return nil
}

func executeHandler(event *types.Event) error {
	log.Println("executing handler with --webHookURL", plugin.WebHookURL)
	return nil
}
