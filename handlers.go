package main

import (
	"fmt"
  "time"
  "strings"
  "github.com/yanzay/tbot/v2"
)

//Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
  msg := fmt.Sprintf("Welcome to Amigolang.\nUse /hello to get Hello.\nUse /say <your word> to say the word, for example /say hello.")
  a.client.SendMessage(m.Chat.ID, msg)
}

//Handle the /hello command here
func (a *application) helloHandler(m *tbot.Message) {
  a.client.SendChatAction(m.Chat.ID, tbot.ActionTyping)
  time.Sleep(1 * time.Second)
  a.client.SendMessage(m.Chat.ID, "Hello!")
}

//Handle the /say command here
func (a *application) sayHandler(m *tbot.Message) {
	text := strings.TrimPrefix(m.Text, "/say ")
	word := fmt.Sprintf("The word you say is: %s", word(text))
	a.client.SendMessage(m.Chat.ID, word)
}

func word(text string) string {
	textLine := fmt.Sprintf(text)
	return textLine
}
