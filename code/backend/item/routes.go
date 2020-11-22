package item

type ItemService interface {
	GetItem(int) (*Item, error)
}

func (s *itemService) GetItem(ID int) (*Item, error) {
	item, err := s.db.getItem(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}
