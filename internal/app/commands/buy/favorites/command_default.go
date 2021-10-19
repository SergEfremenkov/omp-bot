package favorites

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BuyFavoritesCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	c.SendAMessage(inputMessage, "Unknown command. Please use the command /help__buy__favorites for more details!")
}
