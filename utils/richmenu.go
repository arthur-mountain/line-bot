package utils

import "github.com/line/line-bot-sdk-go/v7/linebot"

// Question: Is linebot.RichMenu need pointer
type RichMenu struct {
	*linebot.RichMenu
}

// Question: which is better builder or chain, currently is builder patterns
// Question: can i use my custom struct like this linebot.CreateRichMenu(richmenu),instead of linebot.CreateRichMenu(richmenu.Richmenu)

// Question: Could add default value in the parameter
func CreateRichMenu() RichMenu {
	return RichMenu{}
}

func (r *RichMenu) SetSize(size linebot.RichMenuSize) {
	r.RichMenu.Size = size
}

func (r *RichMenu) SetSelected(selected bool) {
	r.RichMenu.Selected = selected
}

func (r *RichMenu) SetName(name string) {
	r.RichMenu.Name = name
}

func (r *RichMenu) SetChatBarText(chatBarText string) {
	r.RichMenu.ChatBarText = chatBarText
}

//TODO: condition is ares slice or single area
func (r *RichMenu) SetAreas(area linebot.AreaDetail) {
	r.RichMenu.Areas = append(r.Areas, area)
}
