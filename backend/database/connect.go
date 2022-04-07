package database

import (
	"TTU_CORE_ERP_PORTAL/models"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//dsn := "webmaster:PRINT45dull@tcp(localhost:3306)/ucc?charset=utf8&parseTime=True&loc=Local"
	//dsn := "host=localhost user=gadocansey password=1988Gadocansey dbname=UCC port=9920 sslmode=disable TimeZone=Africa/Accra"

	// https://github.com/go-gorm/postgres
	//database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=gadocansey password=1988Gadocansey dbname=UCC port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                 // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		errors.New("error connecting to data")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
