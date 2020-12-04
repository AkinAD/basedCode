package shop

import (
	"log"

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
	getCategories() ([]*Category, error)
	createCategory(string) (*Category, error)
	editCategory(*Category) (*Category, error)
	deleteCategory(int) (bool, error)
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
	var items []*Item
	result := r.db.Raw("select itemid, i.name as name, description, i.categoryid as categoryid, price, c.category as category from items i join categories c on c.categoryid = i.categoryid where itemid = ?", ID).Scan(&items)

	err := result.Error

	if err != nil {
		return nil, err
	}

	log.Printf("[Shop] [Repo] [GetItem] %v | %v", ID, items[0])

	return items[0], nil
}

func (r *shopRepo) addItem(item *Item) (*Item, error) {
	result := r.db.Exec("INSERT INTO items (name, description, categoryid, price) VALUES (?, ?, ?, ?)", item.Name, item.Description, item.CategoryID, item.Price)

	err := result.Error

	if err != nil {
		return nil, err
	}

	return item, nil
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
	result := r.db.Raw("select itemid, i.name as name, description, i.categoryid as categoryid, price, c.category as category from items i join categories c on c.categoryid = i.categoryid").Scan(&items)

	err := result.Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *shopRepo) getItemsFromStore(ID int) ([]*Item, error) {
	return nil, nil
}

func (r *shopRepo) updateItem(item *Item) (*Item, error) {
	// for fields that aren't in request body, it updates item to 0 or empty string. for categoryID, it uses that in WHERE clause instead of updating it
	// result := r.db.Debug().Model(&item).Updates(map[string]interface{}{"name": item.Name, "description": item.Description, "categoryid": item.CategoryID, "price": item.Price})

	result := r.db.Debug().Table("items").Model(&item).Omit("itemid").Updates(&item)

	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *shopRepo) deleteItem(ID int) (bool, error) {
	var items []Item
	result := r.db.Raw("DELETE from items WHERE itemid = ?", ID).Scan(&items)
	err := result.Error

	if err != nil {
		return false, err
	}

	return true, nil
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

func (r *shopRepo) getCategories() ([]*Category, error) {
	var categories []*Category

	result := r.db.Raw("SELECT * FROM categories").Scan(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
func (r *shopRepo) createCategory(name string) (*Category, error) {
	// var cat *Category
	// err := r.db.Exec("INSERT INTO categories (category) VALUES (?) RETURNING *", name).Row().Scan(&cat)
	// if err != nil {
	// 	return nil, err
	// }

	cat := &Category{Name: name}

	result := r.db.Create(&cat)
	if result.Error != nil {
		return nil, result.Error
	}

	return cat, nil
}
func (r *shopRepo) editCategory(cat *Category) (*Category, error) {
	result := r.db.Table("categories").Where("categoryid = ?", cat.CategoryID).Updates(map[string]interface{}{"category": cat.Name})
	if result.Error != nil {
		return nil, result.Error
	}

	return cat, nil
}
func (r *shopRepo) deleteCategory(categoryID int) (bool, error) {
	result := r.db.Table("categories").Delete(&Category{CategoryID: categoryID})
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
