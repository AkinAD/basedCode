package shop

type ShopService interface {
	GetItems() (*Item, error)
	GetItemsFromStore(ID int) (*Item, error)
	GetItem(ID int) (*Item, error)
	AddItem(*Item) (*Item, error)
	UpdateItem(*Item) (*Item, error)
	DeleteItem(int) (bool, error)
	GetStores(ID int) (*Store, error)
	GetStore(ID int) (*Store, error)
	AddStore(*Store) (*Store, error)
	UpdateStore(*Store) (*Store, error)
	DeleteStore(int) (bool, error)
	GetStock(int) (*Stock, error)
	AddStock(int, *Item) (*Stock, error)
	UpdateStock(int, int, *Location) (*Stock, error)
	DeleteAddStock(int, int) (*Stock, error)
}

type shopService struct {
	db ShopRepo
}

func NewService(conn string) ShopService {
	return &shopService{
		db: NewDatabase(conn),
	}
}

type Item struct {
	ID          int `gorm:"<-:false"`
	Name        string
	Description string
	Category    Category
	Price       float64
}

type Store struct {
	ID      int `gorm:"<-:false"`
	Address string
}

type Stock struct {
	StoreID int `gorm:"<-:false"`
	Stock   map[int]Location
}

type Location struct {
	row int
	col int
}

type Category int

const (
	Shirts Category = 1
	Shoes  Category = 2
)

func (s *shopService) GetItems() (*Item, error) {
	item, err := s.db.getItems()
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}
func (s *shopService) GetItemsFromStore(ID int) (*Item, error) {
	item, err := s.db.getItemsFromStore(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) GetItem(ID int) (*Item, error) {
	item, err := s.db.getItem(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) GetStores(ID int) (*Store, error) {
	item, err := s.db.getStores()
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) GetStore(ID int) (*Store, error) {
	item, err := s.db.getStore(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) AddStore(store *Store) (*Store, error) {
	item, err := s.db.addStore(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) UpdateStore(store *Store) (*Store, error) {
	item, err := s.db.updateStore(store)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}
