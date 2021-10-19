package buy

import "fmt"

var FavoritesTestModel = []Favorites{
	{ID: 1, Name: "Name 1", Description: "Description 1"},
	{ID: 2, Name: "Name 2", Description: "Description 2"},
	{ID: 3, Name: "Name 3", Description: "Description 3"},
	{ID: 4, Name: "Name 4", Description: "Description 4"},
}

type Favorites struct {
	ID          uint64
	Name        string
	Description string
}

func (s *Favorites) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Description: %s", s.ID, s.Name, s.Description)
}

func IsIDExists(ID uint64) bool {
	_, err := FindEntityWithID(ID)

	if err != nil {
		return false
	}

	return true
}

func FindEntityWithID(ID uint64) (int, error) {
	for index, entity := range FavoritesTestModel {
		if entity.ID == ID {
			return index, nil
		}
	}

	return 0, fmt.Errorf("entity with id = %d does not exist", ID)
}

/*
func GetNextID() uint64 {
	var maxID uint64

	for _, entity := range FavoritesTestModel {
		if entity.ID > maxID {
			maxID = entity.ID
		}
	}

	maxID++

	return maxID
}
*/
