package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func main() {
	slackAPIToken := os.Getenv("SLACK_API_TOKEN")
	slackChannelID := os.Getenv("SLACK_CHANNEL_ID")

	lambda.Start(func() error {
		slackClient := slack.New(slackAPIToken)

		_, _, err := slackClient.PostMessage(
			slackChannelID,
			slack.MsgOptionText("Hello, world!", false),
		)
	})
}
