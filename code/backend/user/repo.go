package user

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface {
	//getItem(ID int) (*Item, error)
	//add more
	updatePreferredStore(username string, preferredStore int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewDatabase(config string) UserRepo {
	return &userRepo{
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

func (r *userRepo) updatePreferredStore(username string, preferredStore int) error {
	err := r.db.Table("account").Where("username = ?", username).Update("storeID", preferredStore)
	if err != nil {
		return err
	}
	return nil
}
