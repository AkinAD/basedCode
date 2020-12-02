package user

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface {
	updatePreferredStore(username string, preferredStore int) error
}

type userRepo struct {
	db *gorm.DB
}

func newDatabase(config string) UserRepo {
	return &userRepo{
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

func (r *userRepo) updatePreferredStore(username string, preferredStore int) error {
	result := r.db.Table("account").Where("username = ?", username).Update("storeID", preferredStore)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
