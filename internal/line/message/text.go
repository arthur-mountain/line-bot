package message

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// The reply fail message is simple message. actully fail reson will be return and log.
func HandleText(handler *LineMsgHandler, message *linebot.TextMessage) error {
	client := handler.client

	switch message.Text {
	// upload richmenu image
	case "upload":
		imgUrl := filepath.Join(os.Getenv("GOPATH"), "src/line-chatbot/assets/richmenu.png")

		if _, err := client.UploadRichMenuImage(os.Getenv("DEFAULT_RICHMENU_ID"), imgUrl).Do(); err != nil {
			handler.ReplyMessage("upload img fail")

			return err
		}

		handler.ReplyMessage("upload img success")

		//set-default-richmenu
	case "set":
		if _, err := client.SetDefaultRichMenu(os.Getenv("DEFAULT_RICHMENU_ID")).Do(); err != nil {
			handler.ReplyMessage("set default richmenu fail")

			return err
		}

		handler.ReplyMessage("set default richmenu success")

		// cancel-default-richmenu
	case "cancel":
		if _, err := client.CancelDefaultRichMenu().Do(); err != nil {
			handler.ReplyMessage("cancel richmenu fail")

			return err
		}

		handler.ReplyMessage("cancel richmenu success")

		// delete richmenu
	case "delete":
		if _, err := client.DeleteRichMenu(os.Getenv("DEFAULT_RICHMENU_ID")).Do(); err != nil {
			handler.ReplyMessage("delete richmenu fail")

			return err
		}

		handler.ReplyMessage("delete richmenu success")

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
			handler.ReplyMessage("create richmenu fail")

			return err
		}

		handler.ReplyMessage("create richmenu success")

		// get richmenu lists
	case "lists":
		lists, err := client.GetRichMenuList().Do()

		if err != nil {
			handler.ReplyMessage("get richmenu list fail")

			return err
		}

		msg := "get richmenu lists success: \n"

		fmt.Printf("-----------\n length: %d", len(lists))

		for _, list := range lists {
			listJson, _ := json.Marshal(list)

			fmt.Printf("\n listjson: \n %s", listJson)

			msg += fmt.Sprintf(",\n %v", string(listJson))
		}

		handler.ReplyMessage(msg)

		// get default richmenu
	case "get":
		resp, err := client.GetDefaultRichMenu().Do()

		if err != nil {
			handler.ReplyMessage("get default richmenu fail")

			return err
		}

		handler.ReplyMessage(fmt.Sprintf("get default richmenu success: \n %v", resp))

	default:
		msg := fmt.Sprintf("echo message is : \n %s", message.Text)

		handler.ReplyMessage(msg)
	}

	return nil
}
