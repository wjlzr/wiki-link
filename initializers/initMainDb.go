package initializers

import (
	"github.com/spf13/viper"

	db "wiki-link/db"
	"wiki-link/model"
	"wiki-link/model/error"
)

func initMainDb() {
	cfg := viper.Sub("database")
	if cfg == nil {
		panic("config not found database")
	}
	db.InitDatabaseConfig(cfg)
	db.MainDb = db.InitDBConnect()
	db.RecycleDatabaseConfig()
	// CreateDatabase()
	CreateTables()
}

// func CreateDatabase() {
//   db.MainDb.Exec(fmt.Sprintf("create database if not exists %s", db.DatabaseConfig.Name))
// }

func CreateTables() {
	db.MainDb.AutoMigrate(
		&error.Error{},
		&model.Address{},
		&model.BlockHeight{},
		&model.Token{},
	)
}
