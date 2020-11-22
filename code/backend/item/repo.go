package item

import "database/sql"

type ItemRepo interface {
	getItem(ID int) (*Item, error)
	//add more
}

type itemRepo struct {
	db *sql.DB
}

func NewDatabase(config string) ItemRepo {
	return &itemRepo{
		db: newDatabase(config),
	}
}

func newDatabase(config string) *sql.DB {
	// configStr := "user=" + viper.GetString("PGDB.user") +
	// 	" password=" + viper.GetString("PGDB.password") +
	// 	" host=" + viper.GetString("PGDB.host") +
	// 	" port=" + viper.GetString("PGDB.port") +
	// 	" dbname=" + viper.GetString("PGDB.dbname") +
	// 	" sslmode=" + viper.GetString("PGDB.sslmode")

	return nil
}

func (r *itemRepo) getItem(ID int) (*Item, error) {
	return nil, nil
}
