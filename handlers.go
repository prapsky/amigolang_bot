package main

import (
	"fmt"
  "time"
  "strings"
  "github.com/yanzay/tbot/v2"
)

//Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
  msg := "Welcome to Amigolang. Use /hello to greeting."
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
	word := fmt.Sprintf("```\n%s\n```", word(text))
	a.client.SendMessage(m.Chat.ID, word)
}

func word(text string) string {
	textLine := fmt.Sprintf(text)
	return textLine
}
