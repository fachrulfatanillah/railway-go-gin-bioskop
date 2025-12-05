package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=postgres.railway.internal user=postgres password='FTMdCENFpEThEWHGvTwyvebqYZtOAENI' dbname=bioskopdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	DB = db
	fmt.Println("Berhasil konek ke database!")
}
