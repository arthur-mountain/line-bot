package message

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func HandleAudio(handler *LineMsgHandler, message *linebot.StickerMessage) error {
	replyMessage := fmt.Sprintf("sticker id is : %s,\n stickerResourceType is : %s", message.StickerID, message.StickerResourceType)

	handler.ReplyMessage(replyMessage)

	return nil
}
