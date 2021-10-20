package favorites

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
	"log"
)

type ProductData struct {
	ItemID      uint64 `json:"itemID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *BuyFavoritesCommander) New(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.New:"
	newProduct := ProductData{}

	arguments := inputMsg.CommandArguments()

	err := json.Unmarshal([]byte(arguments), &newProduct)
	if err != nil {
		log.Printf("%s error reading json data for type ProductData from input string %v - %v",
			ErrorMessageStart,
			arguments,
			err,
		)

		c.SendAMessage(inputMsg, "Input data error. Use the following syntax:\n\n"+
			"/new__buy__favorites { \"name\" : \"product name\", \"description\": \"product description\" }\n\n"+
			", where field \"name\" is required.",
		)

		return
	}

	itemID, err := c.favoritesService.Create(buy.Favorites(newProduct))
	if err != nil {
		log.Printf("%s %v",
			ErrorMessageStart,
			err,
		)

		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	c.SendAMessage(inputMsg, fmt.Sprintf("Product successfully created with itemID = %d.", itemID))
}
