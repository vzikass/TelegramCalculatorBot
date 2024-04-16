package main

import (
	"fmt"
	"log"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var(
	bot *tgbotapi.BotAPI
	botCommands = []string{"/start", "/help", "/calc"}
)
func main() {
	// You can get a token from BotFather in Telegram.
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Fatal("Error connect to Telegram")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	UpdateConfig := tgbotapi.NewUpdate(0)
	UpdateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(UpdateConfig)
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello! "+
					"This is Telegram bot Calculator. Use \"/help\" to see a list of available commands"))
				bot.Send(msg)
			case "/help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Available commands: %s",
					botCommands))
				bot.Send(msg)
			case "/calc":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("For example:  /calc 2 + 2 . "+
					"You can use: \"+\" \"-\" \"*\"\"/\""))
				bot.Send(msg)

			default:
				//TODO
			}
		}
		// connect calculator
		if strings.HasPrefix(update.Message.Text, "/calc") {
			expression := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/calc"))
			result := calculator(expression)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
