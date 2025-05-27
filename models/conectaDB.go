package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=user password=password dbname=libery port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar com o banco:", err)
    }

    // Migrar as tabelas automaticamente
    database.AutoMigrate(&Users{},&Book{})

    DB = database
}
