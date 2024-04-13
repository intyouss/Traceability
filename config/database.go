package config

import (
	"fmt"

	"github.com/intyouss/Traceability/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.dev") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("db.prefix"),
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("db.maxLifetime"))

	err = db.AutoMigrate(
		&models.User{}, &models.Comment{}, &models.Video{},
		&models.Like{}, &models.Relation{}, &models.Message{},
		&models.Collect{}, &models.MessageOpen{}, &models.UserIncrease{},
		&models.VideoIncrease{}, &models.Role{})
	if err != nil {
		return nil, err
	}
	_ = db.Create(&models.Role{Name: "普通用户", Desc: "普通用户", Status: 1}).Error
	_ = db.Create(&models.Role{Name: "管理员", Desc: "后台管理权限", Status: 1}).Error
	fmt.Println("Database loaded successfully")
	return db, nil
}
