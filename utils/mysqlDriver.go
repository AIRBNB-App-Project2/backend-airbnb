package utils

import (
	"be/configs"
	"be/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)
	fmt.Println(connectionString)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("error in connect database ", err)
		panic(err)
	}

	AutoMigrate(DB)
	return DB
}

func AutoMigrate(DB *gorm.DB) {
	DB.Migrator().DropTable(&entities.User{})
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Room{})
	DB.AutoMigrate(&entities.Image{})
	DB.AutoMigrate(&entities.Booking{})
	DB.AutoMigrate(&entities.Order{})
}
