package favorites

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *BuyFavoritesCommander) generatePaginationButtons(currentOffset uint64) ([]tgbotapi.InlineKeyboardButton, error) {
	paginationButtons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if currentOffset >= c.maxNumOfProductsPerPage {
		button, err := c.createPaginationButton("Previous page", currentOffset-c.maxNumOfProductsPerPage)
		if err != nil {
			return nil, err
		}

		paginationButtons = append(paginationButtons, *button)
	}

	nextOffset := currentOffset + c.maxNumOfProductsPerPage

	_, err := c.favoritesService.List(nextOffset, 1)
	if err == nil {
		button, err := c.createPaginationButton("Next page", nextOffset)
		if err != nil {
			return nil, err
		}

		paginationButtons = append(paginationButtons, *button)
	}

	return paginationButtons, nil
}

func (c *BuyFavoritesCommander) createPaginationButton(buttonName string, offset uint64) (*tgbotapi.InlineKeyboardButton, error) {
	serializedData, err := json.Marshal(CallbackListData{
		Offset: offset,
	})

	if err != nil {
		log.Printf("Error in function generatePaginationButton: "+
			"error encoding data in json - %v",
			err,
		)

		return nil, err
	}

	CallbackPath := path.CallbackPath{
		Domain:       "buy",
		Subdomain:    "favorites",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	button := tgbotapi.NewInlineKeyboardButtonData(buttonName, CallbackPath.String())

	return &button, nil
}
