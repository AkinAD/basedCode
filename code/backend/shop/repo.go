package shop

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShopRepo interface {
	getItem(ID int) (*Item, error)
	addItem(*ItemInStock) (*ItemInStock, error)
	//add more
}

type shopRepo struct {
	db *gorm.DB
}

func NewDatabase(config string) ShopRepo {
	return &shopRepo{
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

func (r *shopRepo) addStock(input *ItemInStock) (*ItemInStock, error) {
	result := r.db.Table("stock").Create(&input)
	if result.Error != nil {
		return nil, err
	}
	return input, nil
}
