package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
  "github.com/yanzay/tbot/v2"
)

type application struct {
  client *tbot.Client
}

var (
  app application
  bot *tbot.Server
  token string
)

func init() {
  e := godotenv.Load()
  if e != nil {
    log.Println(e)
  }
  token = os.Getenv("TOKEN")
}

func main() {
  bot = tbot.New(token, tbot.WithWebhook("https://amigolang-bot.herokuapp.com", ":"+os.Getenv("PORT")))
  app.client = bot.Client()
  bot.HandleMessage("/start", app.startHandler)
  bot.HandleMessage("/hello", app.helloHandler)
  log.Fatal(bot.Start())
}
