package favorites

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
	"log"
)

func (c *BuyFavoritesCommander) Edit(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.Edit:"
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
			"/edit__buy__favorites { \"itemID\" : 1, \"name\" : \"product name\", \"description\": \"product description\" }\n\n"+
			", where the field \"itemID\" is required, its value must be >= 0.",
		)

		return
	}

	err = c.favoritesService.Update(newProduct.ItemID, buy.Favorites(newProduct))
	if err != nil {
		log.Printf("%s %v", ErrorMessageStart, err)
		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	c.SendAMessage(inputMsg, "Product information successfully updated.")
}
