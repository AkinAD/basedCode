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
	result := r.db.First(&items, ID)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println(items[0])
	return &items[0], err
}

func (r *shopRepo) addItem(item *Item) (*Item, error) {
	return nil, nil
}

func (r *shopRepo) addStock(input *StockRequest) (*StockRequest, error) {
	result := r.db.Table("stock").Create(&input)
	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}

func (r *shopRepo) deleteStock(storeID, itemID int) (bool, error) {
	result := r.db.Table("stock").Where("storeID = ? AND itemID = ?", storeID, itemID).Delete(&ItemInStock{})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *shopRepo) getItems() ([]*Item, error) {
	return nil, nil
}

func (r *shopRepo) getItemsFromStore(ID int) ([]*Item, error) {
	return nil, nil
}

func (r *shopRepo) updateItem(item *Item) (*Item, error) {
	return nil, nil
}

func (r *shopRepo) deleteItem(ID int) (bool, error) {
	return false, nil
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
	// result := r.db.Table("stock").Where("storeID = ?", storeID).Find(&itemsInStore)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// itemsInStock := make(map[*Item]*Location)

	// for _, entry := range itemsInStore {
	// 	var item *Item
	// 	result = r.db.Table("items").Where("itemID = ?", entry.ItemID).First(&item)
	// 	if result.Error != nil {
	// 		return nil, result.Error
	// 	}
	// 	itemsInStock[item] = &Location{Row: entry.Row, Col: entry.Col}
	// }

	// stock := &Stock{
	// 	StoreID: storeID,
	// 	Stock:   itemsInStock,
	// }

	stmt := "select i.itemid, i.name as name, description,  c.categoryid, c.name as category, price, row, col from stock join items i on stock.itemid = i.itemid join categories c on c.categoryid = i.categoryid where storeid = ?"
	result := r.db.Raw(stmt, storeID).Scan(&itemsInStore)
	if result.Error != nil {
		return nil, result.Error
	}

	return itemsInStore, nil
}

func (r *shopRepo) addStore(store *Store) (*Store, error) {
	return nil, nil
}

func (r *shopRepo) updateStore(store *Store) (*Store, error) {
	return nil, nil
}

func (r *shopRepo) deleteStore(ID int) (bool, error) {
	return false, nil
}

func (r *shopRepo) updateStock(item *StockRequest) (*StockRequest, error) { return nil, nil }
