package item

type ItemService interface {
	GetItem(int) (*Item, error)
	GetItems() (*[]Item, error)
	DeleteItem(int) (bool, error)
}

func (s *itemService) GetItem(ID int) (*Item, error) {
	item, err := s.db.getItem(ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *itemService) GetItems() (*[]Item, error) {
	item, err := s.db.getItems()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *itemService) DeleteItem(ID int) (bool, error) {
	deleteResult, err := s.db.deleteItem(ID)
	if err != nil {
		return false, err
	}
	return deleteResult, nil
}
