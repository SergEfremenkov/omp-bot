package buy

import "fmt"

var FavoritesTestModel = []Favorites{
	{ItemID: 1, Name: "Name 1", Description: "Description 1"},
	{ItemID: 2, Name: "Name 2", Description: "Description 2"},
	{ItemID: 3, Name: "Name 3", Description: "Description 3"},
	{ItemID: 4, Name: "Name 4", Description: "Description 4"},
	{ItemID: 5, Name: "Name 5", Description: "Description 5"},
	{ItemID: 6, Name: "Name 6", Description: "Description 6"},
	{ItemID: 7, Name: "Name 7", Description: "Description 7"},
	{ItemID: 8, Name: "Name 8", Description: "Description 8"},
	{ItemID: 9, Name: "Name 9", Description: "Description 9"},
}

var SequenceFavoritesTestModel = Sequence{9}

type Favorites struct {
	ItemID      uint64
	Name        string
	Description string
}

func (s *Favorites) String() string {
	return fmt.Sprintf("ItemID: %d, Name: %s, Description: %s", s.ItemID, s.Name, s.Description)
}

func isProductWithIDExists(ID uint64) bool {
	_, err := FindAProductWithID(ID)

	if err != nil {
		return false
	}

	return true
}

func FindAProductWithID(ID uint64) (int, error) {
	for index, product := range FavoritesTestModel {
		if product.ItemID == ID {
			return index, nil
		}
	}

	return 0, fmt.Errorf("product with id = %d does not exist", ID)
}
