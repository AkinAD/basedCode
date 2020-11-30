package item

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ItemRepo interface {
	getItem(ID int) (*Item, error)
	//add more
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
	result := r.db.First(&items, ID)
	err := result.Error

	if err != nil {
		return nil, err
	}
	fmt.Println(items[0])
	return &items[0], err
}
