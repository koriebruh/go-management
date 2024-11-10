package cnf

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/management/domain"
	"log"
	"time"
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

	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ { // Coba ulang hingga 5 kali
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			log.Println("Berhasil terhubung ke database.")
			break
		}

		log.Printf("Percobaan %d: Gagal menghubungkan ke database. Coba lagi dalam 5 detik...\n", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(errors.New("gagal terhubung ke database"))
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
