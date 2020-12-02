package shop

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShopRepo interface {
	getItems() ([]*Item, error)
	getItemsFromStore(ID int) ([]*Item, error)
	getItem(ID int) (*Item, error)
	addItem(*Item) (*Item, error)
	updateItem(*Item) (*Item, error)
	deleteItem(int) (bool, error)
	getStores() ([]*Store, error)
	getStore(ID int) ([]*ItemInStock, error)
	addStore(*Store) (*Store, error)
	updateStore(*Store) (*Store, error)
	deleteStore(int) (bool, error)
	addStock(*StockRequest) (*StockRequest, error)
	updateStock(*StockRequest) (*StockRequest, error)
	deleteStock(int, int) (bool, error)
}

type shopRepo struct {
	db *gorm.DB
}

func newDatabase(config string) ShopRepo {
	return &shopRepo{
		db: initDatabase(config),
	}
}

func initDatabase(config string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (r *shopRepo) getItem(ID int) (*Item, error) {
	var items []Item
	result := r.db.Raw("SELECT * from items WHERE itemid = ?", ID).Scan(&items)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return &items[0], err
}

func (r *shopRepo) addItem(item *Item) (*Item, error) {
	// insert into items (name, description, categoryid, price)
	// values ('Nike React', 'running shoes',2,100.00)
	// var item Item
	result := r.db.Raw("DELETE from items WHERE itemid = ?", nil).Scan(&item)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println("item")
	fmt.Println(item)
	fmt.Println("&item")
	fmt.Println(&item)
	return item, err
}

func (r *shopRepo) addStock(input *StockRequest) (*StockRequest, error) {
	// result := r.db.Table("stock").Create(&input)
	result := r.db.Exec("INSERT INTO stock (storeid, itemid, row, col) VALUES (?, ?, ?, ?)", input.StoreID, input.ItemID, input.Row, input.Col)

	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}

func (r *shopRepo) deleteStock(storeID, itemID int) (bool, error) {
	result := r.db.Exec("DELETE FROM stock WHERE storeID = ? AND itemID = ?", storeID, itemID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *shopRepo) getItems() ([]*Item, error) {
	var items []*Item
	result := r.db.Find(&items)
	fmt.Println("result")
	fmt.Println(result)
	fmt.Println("items")
	fmt.Println(items)

	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return items, err
}

func (r *shopRepo) getItemsFromStore(ID int) ([]*Item, error) {
	return nil, nil
}

func (r *shopRepo) updateItem(item *Item) (*Item, error) {
	result := r.db.Debug().Table("items").Model(&item).Omit("itemid").Updates(&item)

	// result := r.db.Table("items").Where("itemid = ?", item.ItemID).Update("storeID", preferredStore)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("repo.go updateItem")
	fmt.Println(item)
	return item, nil
}

func (r *shopRepo) deleteItem(ID int) (bool, error) {
	var items []Item
	result := r.db.Raw("DELETE from items WHERE itemid = ?", ID).Scan(&items)
	err := result.Error

	if err != nil {
		return false, err
	}
	fmt.Println(items)
	return true, err
}
func (r *shopRepo) getStores() ([]*Store, error) {
	var stores []*Store
	result := r.db.Table("stores").Find(&stores)
	if result.Error != nil {
		return nil, result.Error
	}
	return stores, nil
}

func (r *shopRepo) getStore(storeID int) ([]*ItemInStock, error) {
	var itemsInStore []*ItemInStock
	stmt := "select i.itemid, i.name as name, description,  c.categoryid, c.name as category, price, row, col from stock join items i on stock.itemid = i.itemid join categories c on c.categoryid = i.categoryid where storeid = ?"
	result := r.db.Raw(stmt, storeID).Scan(&itemsInStore)
	if result.Error != nil {
		return nil, result.Error
	}

	return itemsInStore, nil
}

func (r *shopRepo) addStore(store *Store) (*Store, error) {
	result := r.db.Table("stores").Create(&store)
	if result.Error != nil {
		return nil, result.Error
	}

	return store, nil

}

func (r *shopRepo) updateStore(store *Store) (*Store, error) {
	result := r.db.Debug().Table("stores").Model(&store).Omit("storeid").Updates(&store)
	if result.Error != nil {
		return nil, result.Error
	}

	return store, nil
}

func (r *shopRepo) deleteStore(ID int) (bool, error) {
	result := r.db.Table("stores").Delete(&Store{StoreID: ID})
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *shopRepo) updateStock(item *StockRequest) (*StockRequest, error) {
	result := r.db.Table("stock").Model(&item).Omit("storeid", "itemid").Updates(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}
