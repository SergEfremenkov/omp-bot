package favorites

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type FavoritesService interface {
	Describe(favoritesID uint64) (*buy.Favorites, error)
	List(cursor uint64, limit uint64) ([]buy.Favorites, error)
	Create(favorites buy.Favorites) (uint64, error)
	Update(favoritesID uint64, favorites buy.Favorites) error
	Remove(favoritesID uint64) (bool, error)
}

type DummyFavoritesService struct{}

func NewDummyFavoritesService() *DummyFavoritesService {
	return &DummyFavoritesService{}
}

func (s *DummyFavoritesService) Describe(favoritesID uint64) (*buy.Favorites, error) {
	index, err := buy.FindAProductWithID(favoritesID)
	if err != nil {
		return nil, err
	}

	return &buy.FavoritesTestModel[index], nil
}

func (s *DummyFavoritesService) List(cursor uint64, limit uint64) ([]buy.Favorites, error) {
	sliceLength := uint64(len(buy.FavoritesTestModel))
	lastIndex := cursor + limit

	if cursor >= sliceLength {
		return nil, fmt.Errorf("requested boundaries are set incorrectly. Number of products = %d", sliceLength)
	}

	if lastIndex >= sliceLength {
		return buy.FavoritesTestModel[cursor:], nil
	}

	return buy.FavoritesTestModel[cursor:lastIndex], nil
}

func (s *DummyFavoritesService) Create(favorites buy.Favorites) (uint64, error) {
	if favorites.Name == "" {
		return 0, errors.New("the field \"name\" is required")
	}

	favorites.ItemID = buy.SequenceFavoritesTestModel.NextVal()

	buy.FavoritesTestModel = append(buy.FavoritesTestModel, favorites)

	return favorites.ItemID, nil
}

func (s *DummyFavoritesService) Update(favoritesID uint64, favorites buy.Favorites) error {
	index, err := buy.FindAProductWithID(favoritesID)
	if err != nil {
		return err
	}

	buy.FavoritesTestModel[index] = favorites

	return nil
}

func (s *DummyFavoritesService) Remove(favoritesID uint64) (bool, error) {
	index, err := buy.FindAProductWithID(favoritesID)
	if err != nil {
		return false, err
	}

	buy.FavoritesTestModel = append(buy.FavoritesTestModel[:index], buy.FavoritesTestModel[index+1:]...)

	return true, nil
}
