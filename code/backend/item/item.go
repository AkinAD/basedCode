package item

type itemService struct {
	db ItemRepo
}

func NewService(conn string) ItemService {
	return &itemService{
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
	Items   map[Item]Location
}

type Location struct {
	Row int
	Col int
}

type ShoppingCart []Item

type SearchResult []Item

type Category int

const (
	Shirts Category = 1
	Shoes  Category = 2
)
