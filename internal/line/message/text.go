package message

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// The reply fail message is simple message. actully fail reson will be return and log.
func HandleText(event *linebot.Event, client *linebot.Client, message *linebot.TextMessage) error {
	switch message.Text {
	// upload richmenu image
	case "upload":
		imgurl := filepath.Join(os.Getenv("GOPATH"), "src/line-chatbot/assets/richmenu.png")

		if _, err := client.UploadRichMenuImage(os.Getenv("DEFAULT_RICHMENU_ID"), imgurl).Do(); err != nil {
			replyMessage(client, event.ReplyToken, "upload img fail")

			return err
		}

		replyMessage(client, event.ReplyToken, "upload img success")

		//set-default-richmenu
	case "set":
		if _, err := client.SetDefaultRichMenu(os.Getenv("DEFAULT_RICHMENU_ID")).Do(); err != nil {
			replyMessage(client, event.ReplyToken, "set default richmenu fail")

			return err
		}

		replyMessage(client, event.ReplyToken, "set default richmenu success")

		// cancel-default-richmenu
	case "cancel":
		if _, err := client.CancelDefaultRichMenu().Do(); err != nil {
			replyMessage(client, event.ReplyToken, "cancel richmenu fail")

			return err
		}

		replyMessage(client, event.ReplyToken, "cancel richmenu success")

		// delete richmenu
	case "delete":
		if _, err := client.DeleteRichMenu(os.Getenv("DEFAULT_RICHMENU_ID")).Do(); err != nil {
			replyMessage(client, event.ReplyToken, "delete richmenu fail")

			return err
		}

		replyMessage(client, event.ReplyToken, "delete richmenu success")

		// create new richmenu
	case "create":
		richmenu := linebot.RichMenu{
			Size:        linebot.RichMenuSize{Width: 2500, Height: 843},
			Selected:    false,
			Name:        "NARI RICHMENU",
			ChatBarText: "NARI",
			Areas: []linebot.AreaDetail{
				{
					Bounds: linebot.RichMenuBounds{X: 0, Y: 0, Width: 833, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "set", // Will invoke 'set' keyword to set default richmenu
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 834, Y: 0, Width: 833, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeMessage,
						Text: "TODO: Wait liff website then change to liff url",
					},
				},
				{
					Bounds: linebot.RichMenuBounds{X: 1668, Y: 0, Width: 833, Height: 843},
					Action: linebot.RichMenuAction{
						Type: linebot.RichMenuActionTypeURI,
						URI:  "https://zh-hant.reactjs.org/",
						Text: "React",
					},
				},
			},
		}

		if _, err := client.CreateRichMenu(richmenu).Do(); err != nil {
			replyMessage(client, event.ReplyToken, "create richmenu fail")

			return err
		}

		replyMessage(client, event.ReplyToken, "create richmenu success")

		// get richmenu lists
	case "lists":
		lists, err := client.GetRichMenuList().Do()

		if err != nil {
			replyMessage(client, event.ReplyToken, "get richmenu list fail")

			return err
		}

		msg := "get richmenu lists success: \n"

		fmt.Printf("-----------\n length: %d", len(lists))

		for _, list := range lists {
			listJson, _ := json.Marshal(list)

			fmt.Printf("\n listjson: \n %s", listJson)

			msg += fmt.Sprintf(",\n %v", string(listJson))
		}

		replyMessage(client, event.ReplyToken, msg)

		// get default richmenu
	case "get":
		resp, err := client.GetDefaultRichMenu().Do()

		if err != nil {
			replyMessage(client, event.ReplyToken, "get default richmenu fail")

			return err
		}

		replyMessage(client, event.ReplyToken, fmt.Sprintf("get default richmenu success: \n %v", resp))

	default:
		msg := fmt.Sprintf("echo message is : \n %s", message.Text)

		replyMessage(client, event.ReplyToken, msg)
	}

	return nil
}

func replyMessage(client *linebot.Client, replyToken, message string) {
	_, err := client.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do()

	if err != nil {
		fmt.Printf("reply message fail: %v", err.Error())
	}
}
