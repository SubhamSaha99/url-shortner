package database

import (
	"fmt"

	"github.com/SubhamSaha99/url-shortner/internal/modules/url/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
    db, err := gorm.Open(
        sqlite.Open("url-shortener.db"),
        &gorm.Config{},
    )
    if err != nil {
        return nil, fmt.Errorf("connect to SQLite: %w", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("get database connection: %w", err)
    }

    if err := sqlDB.Ping(); err != nil {
        _ = sqlDB.Close()
        return nil, fmt.Errorf("ping SQLite: %w", err)
    }

    if err := db.AutoMigrate(&models.URL{}); err != nil {
        _ = sqlDB.Close()
        return nil, fmt.Errorf("migrate URL model: %w", err)
    }

    return db, nil
}