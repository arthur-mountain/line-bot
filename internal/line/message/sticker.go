package message

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HandleSticker(event *linebot.Event, client *linebot.Client, message *linebot.StickerMessage) error {
	replyMessage := fmt.Sprintf("sticker id is : %s,\n stickerResourceType is : %s", message.StickerID, message.StickerResourceType)

	_, err := client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()

	return err
}
