package db

import (
	"fmt"
	"log"

	"gitlab.com/veeery/gdi_echo_golang.git/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		configuration.DB_USERNAME, 
		configuration.DB_PASSWORD, 
		configuration.DB_HOST, 
		configuration.DB_PORT ,
		configuration.DB_NAME)

	gormDB, err  = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	

	for _, model := range MigrateTable() {
		err := gormDB.Debug().AutoMigrate(model.Table)

		if err != nil {
			log.Fatal(err)
		}
	}

}

func DbManager() *gorm.DB {
	return gormDB
}