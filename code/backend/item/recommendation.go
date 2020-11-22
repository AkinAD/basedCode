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
	ID       int
	Name     string
	Category Category
	Price    int
}

type Store struct {
	ID        int
	Name      string
	Address   string //?
	Latitude  int
	Longitude int
	Items     map[Item]Location
}

type Location struct {
	Row int
	Col int
}

type ShoppingCart []Item

type SearchResult []Item

type Category string

const (
	CAT1 Category = "CAT1"
	CAT2 Category = "CAT2"
)
