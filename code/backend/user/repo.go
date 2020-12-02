package user

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface {
	updatePreferredStore(username string, preferredStore int) error
	updateProfile(input *User) (*User, error)
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
	result := r.db.Table("accounts").Where("username = ?", username).Update("storeID", preferredStore)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepo) updateProfile(input *User) (*User, error) {
	result := r.db.Table("accounts").Where("username = ?", input.Username).Update("storeid", input.StoreID)
	//result := r.db.Model(&User{StoreID: input.StoreID})
	//var users []User
	//result := r.db.Raw("update accounts set storeid = 1 where username = 'matthewwalk';").Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}
