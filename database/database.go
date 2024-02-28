package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ozanpay/model"
)

const (
	HOST     = "localhost"
	DATABASE = "ozanpay"
	USER     = "ozanpay"
	PASSWORD = "ozanpay2024!!!"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Connect() {
	vt := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
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
func ConnectAndMigrate() {
	Connect()
	Migrate()
}
