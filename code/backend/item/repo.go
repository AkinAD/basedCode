package item

import (
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
	return nil, nil
}
