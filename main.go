package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Replace TOKEN with your bot's token from BotFather
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal("Error connect to Telegram")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set commands for the bot
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Start interacting with the bot"},
		{Command: "help", Description: "Show available commands"},
		{Command: "calc", Description: "Perform a calculation (e.g., /calc 2 + 2)"},
	}

	setCommands := tgbotapi.NewSetMyCommands(commands...)
	_, err = bot.Request(setCommands)
	if err != nil {
		log.Fatal("Failed to set bot commands:", err)
	}
	// Configuration to receive updates
	UpdateConfig := tgbotapi.NewUpdate(0)
	UpdateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(UpdateConfig)

	// Handling updates from the bot
	for update := range updates {
		if update.CallbackQuery != nil {
			// Handle callback query
			callback := update.CallbackQuery
			if callback.Data == "calc_pressed" {
				// Send message asking for expression
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Please enter the expression (e.g., /calc 2 + 2). You can use +, -, *, / or ! after the number. Example: 5!")
				bot.Send(msg)

				// Optionally, acknowledge the callback
				callbackMsg := tgbotapi.NewCallback(callback.ID, "Instruction sent")
				bot.Send(callbackMsg)
			} else if update.Message != nil {
				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello! "+
						"This is Telegram bot Calculator. Use \"/help\" to see a list of available commands"))
					bot.Send(msg)

				case "/help":
					helpText := "Available commands:\n" +
						"/start - Start interacting with the bot\n" +
						"/help - Show available commands\n" +
						"/calc - Perform a calculation (e.g., /calc 2 + 2)"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)

					keyboard := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("Calculate", "calc_pressed"),
						),
					)
					msg.ReplyMarkup = keyboard
					bot.Send(msg)

				case "/calc":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("For example:  /calc 2 + 2 . "+
						"You can use: + - * / or ! after the num, example: 5!"))
					bot.Send(msg)
				}
			}

			// Handling clicks on the Inline button
			if update.CallbackQuery != nil {
				// Handle callback query
				callback := update.CallbackQuery
				if callback.Data == "calc_pressed" {
					// Send message asking for expression
					msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Please enter the expression (e.g., /calc 2 + 2)")
					bot.Send(msg)

					// Optionally, acknowledge the callback
					callbackMsg := tgbotapi.NewCallback(callback.ID, "You can now enter your calculation")
					bot.Send(callbackMsg)
				}

				// connect calculator
				if strings.HasPrefix(update.Message.Text, "/calc") {
					expression := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/calc"))
					result := calculator(expression)
					log.Printf("Calculation result: %s", result)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
					msg.ReplyToMessageID = update.Message.MessageID
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Failed to send message: %v", err)
					}
				}
			}
		}
	}
}
