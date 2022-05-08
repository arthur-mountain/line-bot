package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"line-chatbot/internal/line/message"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	client *linebot.Client
	err    error
)

func main() {
	client, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))

	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/callback", callbackHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := client.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		// fmt.Printf("event : \n %+v \n", event)

		if event.Type == linebot.EventTypeMessage {
			fmt.Printf("message event : \n %#v \n", event.Message)

			msgHandler := message.NewMsgHandler(event, client)

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msgHandler.Text(message)

			case *linebot.StickerMessage:
				msgHandler.Sticker(message)
			}
		}

		if event.Type == linebot.EventTypePostback {
			fmt.Println("TODO: handle postback")
		}
	}
}
