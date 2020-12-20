package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, _ := tgbotapi.NewBotAPI("BotToken")
	bot.Debug = true
	var ChatID int64
	ChatID = 1316224975 //Chatid

	for i := 1; i < len(os.Args); i++ {
		msg := tgbotapi.NewDocumentUpload(ChatID, os.Args[i])
		bot.Send(msg)
	}

}
