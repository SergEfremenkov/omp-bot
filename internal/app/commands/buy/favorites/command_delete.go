package favorites

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *BuyFavoritesCommander) Delete(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.Delete:"

	arguments := inputMsg.CommandArguments()

	ID, err := strconv.Atoi(arguments)
	if err != nil || ID <= 0 {
		msg := fmt.Sprintf("wrong argument. Entity ID must be greater than 0")

		log.Printf("%s %s\nCommandArguments = %s", ErrorMessageStart, msg, arguments)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %s", msg))

		return
	}

	_, err = c.favoritesService.Remove(uint64(ID))
	if err != nil {
		log.Printf("%s %v", ErrorMessageStart, err)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	c.SendAMessage(inputMsg, "Entity successfully deleted")
}
