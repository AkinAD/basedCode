package shop

type ShopService interface {
	GetItems() ([]*Item, error)
	GetItemsFromStore(ID int) ([]*Item, error)
	GetItem(ID int) (*Item, error)
	CreateItem(*Item) (*Item, error)
	UpdateItem(*Item) (*Item, error)
	DeleteItem(int) (bool, error)
	GetStores() ([]*Store, error)
	GetStore(ID int) ([]*ItemInStock, error)
	CreateStore(*Store) (*Store, error)
	UpdateStore(*Store) (*Store, error)
	DeleteStore(int) (bool, error)
	CreateStock(*StockRequest) (*StockRequest, error)
	UpdateStock(*StockRequest) (*StockRequest, error)
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
	ItemID      int     `json:"itemID" gorm:"->;primaryKey;column:itemid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category
}

type Store struct {
	StoreID int    `json:"storeID" gorm:"->;primaryKey;column:storeid"`
	Address string `json:"address"`
}

type ItemInStock struct {
	Item
	Location
}

type StockRequest struct {
	StoreID int `json:"storeID" gorm:"->;primaryKey;column:storeid"`
	ItemID  int `json:"itemID" gorm:"->;primaryKey;column:itemid"`
	Location
}

type Location struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Category struct {
	CategoryID int    `json:"categoryID" gorm:"->;primaryKey;column:categoryid"`
	Name       string `json:"category" gorm:"column:category"`
}

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

func (s *shopService) CreateItem(item *Item) (*Item, error) {
	item, err := s.db.addItem(item)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) GetStores() ([]*Store, error) {
	stores, err := s.db.getStores()
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return stores, nil
}

func (s *shopService) GetStore(ID int) ([]*ItemInStock, error) {
	stock, err := s.db.getStore(ID)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return stock, nil
}

func (s *shopService) CreateStore(store *Store) (*Store, error) {
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

func (s *shopService) UpdateItem(item *Item) (*Item, error) {
	return nil, nil
}
func (s *shopService) DeleteItem(itemID int) (bool, error) {
	deleteResult, err := s.db.deleteItem(ID)
	if err != nil {
		return false, err
	}
	return deleteResult, nil
}
func (s *shopService) DeleteStore(storeID int) (bool, error) {
	result, err := s.db.deleteStore(storeID)
	if err != nil {
		// log.Printf("%v", err)
		return false, err
	}
	return result, nil
}
func (s *shopService) CreateStock(request *StockRequest) (*StockRequest, error) {
	item, err := s.db.addStock(request)
	if err != nil {
		// log.Printf("%v", err)
		return nil, err
	}

	return item, nil
}

func (s *shopService) UpdateStock(request *StockRequest) (*StockRequest, error) {
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
