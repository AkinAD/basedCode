package user

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface {
	updatePreferredStore(username string, preferredStore int) error
	updateProfile(input *User) (*User, error)
	getProfile(input string) (*User, error)
	createProfile(Username string, StoreID int, FirstName string, LastName string) error
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

func (r *userRepo) getProfile(input string) (*User, error) {
	var userProfile User
	result := r.db.Raw("SELECT * from accounts WHERE username = ?", input).Scan(&userProfile)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userProfile, nil
}
func (r *userRepo) createProfile(Username string, StoreID int, FirstName string, LastName string) error {
	//result := r.db.Table("accounts").Where("username = ?", input.Username).Update("storeid", input.StoreID)
	result := r.db.Exec("INSERT INTO accounts (username,storeid,firstname,lastname) VALUES (?,?,?,?)", Username, StoreID, FirstName, LastName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
