package message

import (
	"fmt"
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
	if err := HandleText(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle sticker */
func (handler *LineMsgHandler) Sticker(message *linebot.StickerMessage) {
	if err := HandleSticker(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle image */
func (handler *LineMsgHandler) Image(message *linebot.StickerMessage) {
	if err := HandleImage(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle video */
func (handler *LineMsgHandler) Video(message *linebot.StickerMessage) {
	if err := HandleImage(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle location */
func (handler *LineMsgHandler) Location(message *linebot.StickerMessage) {
	if err := HandleImage(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* Handle audio */
func (handler *LineMsgHandler) Audio(message *linebot.StickerMessage) {
	if err := HandleImage(handler, message); err != nil {
		log.Println(err.Error())
	}
}

/* common usage */
func (handler *LineMsgHandler) ReplyMessage(message string) {
	_, err := handler.client.ReplyMessage(handler.event.ReplyToken, linebot.NewTextMessage(message)).Do()

	if err != nil {
		fmt.Printf("reply message fail: %v", err.Error())
	}
}
