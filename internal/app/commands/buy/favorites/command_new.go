package favorites

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/buy"
	"log"
)

type EntityData struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *BuyFavoritesCommander) New(inputMsg *tgbotapi.Message) {
	const ErrorMessageStart = "Error in function BuyFavoritesCommander.New:"
	newEntity := EntityData{}

	arguments := inputMsg.CommandArguments()

	err := json.Unmarshal([]byte(arguments), &newEntity)
	if err != nil {
		log.Printf("%s error reading json data for type EntityData from input string %v - %v",
			ErrorMessageStart,
			arguments,
			err,
		)

		c.SendAMessage(inputMsg, "Input data error. Use the following syntax:\n\n"+
			"/new__buy__favorites { \"name\" : \"entity name\", \"description\": \"entity description\" }\n\n"+
			", where field \"name\" is required.",
		)

		return
	}

	ID, err := c.favoritesService.Create(buy.Favorites(newEntity))
	if err != nil {
		log.Printf("%s %v",
			ErrorMessageStart,
			err,
		)

		c.SendAMessage(inputMsg, fmt.Sprintf("Error: %v", err))

		return
	}

	c.SendAMessage(inputMsg, fmt.Sprintf("Entity successfully created with ID = %d.", ID))
}
