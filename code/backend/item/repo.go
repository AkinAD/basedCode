package item

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ItemRepo interface {
	getItem(ID int) (*Item, error)
	getItems() (*[]Item, error)
	deleteItem(ID int) (bool, error)
	addItem(itemToAdd Item) (*Item, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewDatabase(config string) ItemRepo {
	return &itemRepo{
		db: newDatabase(config),
	}
}

func newDatabase(config string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (r *itemRepo) getItem(ID int) (*Item, error) {
	var items []Item
	result := r.db.Raw("SELECT * from items WHERE itemid = ?", ID).Scan(&items)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return &items[0], err
}

func (r *itemRepo) getItems() (*[]Item, error) {
	var items []Item
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
	return &items, err
}

func (r *itemRepo) deleteItem(ID int) (bool, error) {
	var items []Item
	result := r.db.Raw("DELETE from items WHERE itemid = ?", ID).Scan(&items)
	err := result.Error

	if err != nil {
		return false, err
	}
	fmt.Println(items)
	return true, err
}

//INSERT INTO public.items (itemid, name, description, categoryid, price) VALUES (4, 'Crocs', 'wear this for WAP', 2, 30.00)

func (r *itemRepo) addItem(itemToAdd Item) (*Item, error) {
	// insert into items (name, description, categoryid, price)
	// values ('Nike React', 'running shoes',2,100.00)
	var item Item
	result := r.db.Raw("DELETE from items WHERE itemid = ?", nil).Scan(&item)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println("item")
	fmt.Println(item)
	fmt.Println("&item")
	fmt.Println(&item)
	return &item, err
}

// allen says .Raw() is for queries & .Exec() is for stuff like update, insert, delete
