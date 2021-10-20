package favorites

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BuyFavoritesCommander) List(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.List:"
	var initialOffset uint64 = 0

	listOfProducts, err := c.favoritesService.List(initialOffset, c.maxNumOfProductsPerPage)
	if err != nil {
		log.Printf("%s %v", ErrorMessageStart, err)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	paginationButtons, err := c.generatePaginationButtons(initialOffset)
	if err != nil {
		log.Println(err)
		return
	}

	var outputMsgText string = "List of products:\n\n"
	for _, product := range listOfProducts {
		outputMsgText += product.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		outputMsgText,
	)

	if len(paginationButtons) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				paginationButtons...,
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("%s error sending reply message to chat - %v", ErrorMessageStart, err)
	}
}
