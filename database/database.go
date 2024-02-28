package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ozanpay/config"
	"ozanpay/model"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Connect(cfg config.DbConfig) {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name)
	var err error
	db, err = gorm.Open(postgres.Open(vt), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
}

func Migrate() {
	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.Seller{})
	db.AutoMigrate(model.Organization{})
}
func ConnectAndMigrate(cfg config.DbConfig) {
	Connect(cfg)
	Migrate()
}
