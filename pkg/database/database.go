package database

//https://gorm.io/ru_RU/docs/

import (
	"DababaseService/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var dbase *gorm.DB
var g_cfg *config.Config

// Init - Инициализация базы данных
func Init(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Database.Host,
		cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	g_cfg = cfg

	return db, nil
}

// GetDB - Получение ссылки на экземпляр базы данных
func GetDB() *gorm.DB {
	if dbase == nil {
		dbase, _ = Init(g_cfg)
		sleep := time.Duration(1)
		for dbase == nil {
			sleep += 1
			fmt.Printf("Не удалось подключиться к базе данных, повторное подключение через %d секунд", sleep)
			time.Sleep(sleep * time.Second)
			dbase, _ = Init(g_cfg)
		}
	}
	return dbase
}

func CloseConnection() {
	db := GetDB()
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
	fmt.Println("Close database connections")
}
