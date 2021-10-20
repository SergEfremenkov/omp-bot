package favorites

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/buy/favorites"
	"log"
)

type FavoritesCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BuyFavoritesCommander struct {
	bot                     *tgbotapi.BotAPI
	favoritesService        favorites.FavoritesService
	maxNumOfProductsPerPage uint64
}

func NewFavoritesCommander(bot *tgbotapi.BotAPI) FavoritesCommander {
	favoritesService := favorites.NewDummyFavoritesService()

	return &BuyFavoritesCommander{
		bot:                     bot,
		favoritesService:        favoritesService,
		maxNumOfProductsPerPage: 4,
	}
}

func (c *BuyFavoritesCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BuyFavoritesCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BuyFavoritesCommander) HandleCommand(inputMsg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(inputMsg)
	case "get":
		c.Get(inputMsg)
	case "list":
		c.List(inputMsg)
	case "delete":
		c.Delete(inputMsg)
	case "new":
		c.New(inputMsg)
	case "edit":
		c.Edit(inputMsg)
	default:
		c.Default(inputMsg)
	}
}

func (c *BuyFavoritesCommander) SendAMessage(inputMsg *tgbotapi.Message, messageToSend string) bool {
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		messageToSend,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BuyFavoritesCommander.SendAMessage: error sending reply message to chat - %v", err)
		return false
	}

	return true
}
