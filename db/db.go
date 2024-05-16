package db

import (
	"errors"
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"github.com/AlchemillaHQ/gotham-stack/internal/config"
	"github.com/AlchemillaHQ/gotham-stack/internal/logger"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"path"
)

func SetupDatabase(env string, dbURL string, dataPath string) *gorm.DB {
	var db *gorm.DB
	var err error

	ormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	if env == "production" || env == "prod" {
		db, err = gorm.Open(postgres.Open(dbURL), ormConfig)
	} else {
		db, err = gorm.Open(sqlite.Open(path.Join(dataPath, "gotham-stack.db")), ormConfig)
	}

	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		logger.Error("failed to migrate database", zap.Error(err))
	}

	return db
}

func SetupBasicData(db *gorm.DB, admins []config.Admin) {
	for _, admin := range admins {
		var user models.User
		err := db.Where("username = ?", admin.Username).Or("email = ?", admin.Email).First(&user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, hashErr := utils.HashPassword(admin.Password)
			if hashErr != nil {
				logger.Error("Failed to hash password for admin", zap.String("username", admin.Username), zap.Error(hashErr))
				continue
			}

			newUser := models.User{
				Username: admin.Username,
				Password: hashedPassword,
				Email:    admin.Email,
			}
			result := db.Create(&newUser)
			if result.Error != nil {
				logger.Error("Failed to create admin user", zap.String("username", admin.Username), zap.Error(result.Error))
				continue
			}
		} else if err != nil {
			log.Printf("Error checking for admin user %v: %v", admin.Username, err)
		}
	}

	logger.Info("Database initialized")
}
