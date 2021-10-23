package buy

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/buy/favorites"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BuyCommander struct {
	bot                *tgbotapi.BotAPI
	favoritesCommander Commander
}

func NewBuyCommander(
	bot *tgbotapi.BotAPI,
) *BuyCommander {
	return &BuyCommander{
		bot:                bot,
		favoritesCommander: favorites.NewFavoritesCommander(bot),
	}
}

func (c *BuyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "favorites":
		c.favoritesCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BuyCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BuyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "favorites":
		c.favoritesCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BuyCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
