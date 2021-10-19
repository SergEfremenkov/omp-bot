package favorites

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BuyFavoritesCommander) Help(inputMsg *tgbotapi.Message) {
	messageToSend := "/help__buy__favorites - print list of commands\n" +
		"/get__buy__favorites - get a entity\n" +
		"/list__buy__favorites - get a list of your entity\n" +
		"/delete__buy__favorites - delete an existing entity\n" +
		"/new__buy__favorites - create a new entity\n" +
		"/edit__buy__favorites - edit a entity"

	c.SendAMessage(inputMsg, messageToSend)
}
