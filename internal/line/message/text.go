package message

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HandleText(event *linebot.Event, client *linebot.Client, message *linebot.TextMessage) error {
	replyMessage := fmt.Sprintf("echo message is : \n %s", message.Text)
	_, err := client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()

	return err
}
