package favorites

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *BuyFavoritesCommander) Get(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.Get:"

	arguments := inputMsg.CommandArguments()

	itemID, err := strconv.Atoi(arguments)
	if err != nil || itemID < 0 {
		msg := fmt.Sprintf("wrong argument. Product ID must be an unsigned integer.")

		log.Printf("%s %s\nCommandArguments = %s", ErrorMessageStart, msg, arguments)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %s", msg))

		return
	}

	product, err := c.favoritesService.Describe(uint64(itemID))
	if err != nil {
		log.Printf("%s %v", ErrorMessageStart, err)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	c.SendAMessage(inputMsg, product.String())
}
