package main

import (
  "time"
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
