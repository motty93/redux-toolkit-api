package main

import (
	"app/config"
	"app/pkg/db/model"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func seeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		task := model.Task{Title: "title" + strconv.Itoa(i)}
		if err := db.Create(&task).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func init() {
	config.Init()
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := config.Config.DbUserName + ":" + config.Config.DbPassword + "@tcp(" + config.Config.DbHost + ":" + config.Config.DbPort + ")/" + config.Config.DbName + "?" + config.Config.DbParams
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error(err)
	}

	if err := seeds(DB); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
