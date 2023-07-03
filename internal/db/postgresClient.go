package db

import (
	"fmt"
	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type PostgresClient struct {
	Db         *gorm.DB
	InfoLogger *log.Logger
}

func InitPostgresClient(loggers logger.Loggers) (postgresClient PostgresClient, err error) {
	// TODO set stdout gorm logger depends on app mode
	fmt.Println("init postgres")
	gormLogger := gormLogger.New(
		log.New(os.Stdout, "[GORM]\t", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres_dsn")), &gorm.Config{Logger: gormLogger})
	if err != nil {
		loggers.Err.Fatalf("Failed to initialize postgres client: %s", err.Error())
		return
	}
	postgresClient = PostgresClient{
		Db:         db,
		InfoLogger: loggers.Info,
	}
	if migrateErr := postgresClient.Migrate(); migrateErr != nil {
		loggers.Err.Fatalf("Failed to migrate: %s", migrateErr.Error())
	}
	return
}

func (c *PostgresClient) Migrate() (err error) {
	err = c.Db.AutoMigrate(
		&models.UserCore{},
		&models.ProjectPageCore{},
		&models.ProjectCore{},
		&models.ParentRelCore{},
		&models.SettingsCore{},
	)
	if err != nil {
		return err
	}
	var count int64
	if err := c.Db.First(&models.SettingsCore{ID: 1}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	} else {
		return c.Db.Create(&models.SettingsCore{ID: 1, ActivationByLink: true}).Error
	}
	return nil
}
