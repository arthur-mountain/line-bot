package message

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineMsgHandler struct {
	event  *linebot.Event
	client *linebot.Client
}

func NewMsgHandler(event *linebot.Event, client *linebot.Client) LineMsgHandler {
	return LineMsgHandler{event, client}
}

/* Handle text */
func (handler *LineMsgHandler) Text(message *linebot.TextMessage) {
	if err := HandleText(handler.event, handler.client, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle sticker */
func (handler *LineMsgHandler) Sticker(message *linebot.StickerMessage) {
	if err := HandleSticker(handler.event, handler.client, message); err != nil {
		log.Println(err.Error())
	}
}
