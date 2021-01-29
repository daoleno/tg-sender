package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := flag.String("token", "", "bot token")
	chainID := flag.Int64("chat-id", 0, "chat id")
	fileName := flag.String("file", "", "file name")
	flag.Parse()

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewDocumentUpload(*chainID, fileName)
	m, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	log.Println(m)
}
