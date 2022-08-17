package service

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"twodo.app/condo/model"
)

type dbService struct {
	Create  func() (*gorm.DB, error)
	Migrate func()
}

var config = model.LoadConfig()
var databaseConfig = config.Database
var dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", databaseConfig.Host, databaseConfig.User, databaseConfig.Password, databaseConfig.DbName, databaseConfig.Port)

var DB = dbService{
	Create: func() (*gorm.DB, error) {
		db, err := createDB()
		return db, err
	},
	Migrate: func() {
		db, err := createDB()
		if err != nil {
			log.Fatal(err)
		}

		if err := db.AutoMigrate(&model.WaterBillInfo{}); err != nil {
			log.Fatal(err)
		}
	},
}

func createDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
