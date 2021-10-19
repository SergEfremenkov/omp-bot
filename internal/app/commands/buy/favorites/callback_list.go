package favorites

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *BuyFavoritesCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.CallbackList:"
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("%s error reading json data for type CallbackListData from input string %v - %v",
			ErrorMessageStart,
			callbackPath.CallbackData,
			err,
		)

		return
	}

	listOfEntities, err := c.favoritesService.List(parsedData.Offset, c.maxNumOfEntitiesPerPage)
	if err != nil {
		log.Printf("%s %v", ErrorMessageStart, err)
		c.SendAMessage(callback.Message, fmt.Sprintf("Error: %v", err))

		return
	}

	paginationButtons, err := c.generatePaginationButtons(parsedData.Offset)
	if err != nil {
		log.Println(err)
		return
	}

	outputMsgText := "Entity List:\n\n"
	for _, entity := range listOfEntities {
		outputMsgText += entity.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
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
