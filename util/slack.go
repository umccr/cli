package util

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func SendMessage() {
	//viper.AddConfigPath(util.sys.findHome())
	viper.SetConfigName(".umccr")
	viper.ReadInConfig()

	api := slack.New(viper.GetString("SLACK_TOKEN"))
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage("CHANNEL_ID", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
