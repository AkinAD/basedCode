package shop

type ShopService interface {
	GetItems() ([]*Item, error)
	GetItemsFromStore(ID int) ([]*Item, error)
	GetItem(ID int) (*Item, error)
	AddItem(*Item) (*Item, error)
	UpdateItem(*Item) (*Item, error)
	DeleteItem(int) (bool, error)
	GetStores() (*Store, error)
	GetStore(ID int) (*Store, error)
	AddStore(*Store) (*Store, error)
	UpdateStore(*Store) (*Store, error)
	DeleteStore(int) (bool, error)
	GetStock(int) (*Stock, error)
	AddStock(*ItemInStock) (*ItemInStock, error)
	UpdateStock(*ItemInStock) (*ItemInStock, error)
	DeleteStock(int, int) (bool, error)
}

type shopService struct {
	db ShopRepo
}

func NewService(conn string) ShopService {
	return &shopService{
		db: newDatabase(conn),
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

type ItemInStock struct {
	StoreID int `json:"storeID"`
	ItemID  int `json:"itemID"`
	Row     int `json:"row"`
	Col     int `json:"col"`
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

func (s *shopService) GetItems() ([]*Item, error) {
	item, err := s.db.getItems()
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}
func (s *shopService) GetItemsFromStore(ID int) ([]*Item, error) {
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

func (s *shopService) AddItem(item *Item) (*Item, error) {
	item, err := s.db.addItem(item)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) GetStores() (*Store, error) {
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
	item, err := s.db.addStore(store)
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

func (s *shopService) UpdateItem(*Item) (*Item, error) {
	return nil, nil
}
func (s *shopService) DeleteItem(int) (bool, error) {
	return false, nil
}
func (s *shopService) DeleteStore(int) (bool, error) {
	return false, nil
}
func (s *shopService) GetStock(int) (*Stock, error) {
	return nil, nil
}
func (s *shopService) AddStock(request *ItemInStock) (*ItemInStock, error) {
	item, err := s.db.addStock(request)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) UpdateStock(request *ItemInStock) (*ItemInStock, error) {
	item, err := s.db.updateStock(request)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}
func (s *shopService) DeleteStock(shopID int, itemID int) (bool, error) {
	result, err := s.db.deleteStock(shopID, itemID)
	if err != nil {
		// log.Printf("%v", err)
		return false, err
	}
	return result, nil
}
