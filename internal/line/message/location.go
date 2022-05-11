package message

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HandleLocation(handler *LineMsgHandler, message *linebot.StickerMessage) error {
	client := handler.client
	event := handler.event

	replyMessage := fmt.Sprintf("sticker id is : %s,\n stickerResourceType is : %s", message.StickerID, message.StickerResourceType)

	_, err := client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()

	return err
}
