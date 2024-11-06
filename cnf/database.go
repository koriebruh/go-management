package cnf

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/management/domain"
)

func InitDB() *gorm.DB {
	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Pass,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(errors.New("Failed Connected into database"))
	}

	// Auto Migrate
	err = db.AutoMigrate(
		&domain.Admin{},
		&domain.Category{},
		&domain.Supplier{},
		&domain.Item{},
	)
	if err != nil {
		panic(errors.New("Failed Migrated"))
	}

	return db
}
