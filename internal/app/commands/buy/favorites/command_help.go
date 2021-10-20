package favorites

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BuyFavoritesCommander) Help(inputMsg *tgbotapi.Message) {
	messageToSend := "/help__buy__favorites - print list of commands\n" +
		"/get__buy__favorites - get a product\n" +
		"/list__buy__favorites - get a list of products\n" +
		"/delete__buy__favorites - delete an existing product\n" +
		"/new__buy__favorites - create a new product\n" +
		"/edit__buy__favorites - edit a product"

	c.SendAMessage(inputMsg, messageToSend)
}
