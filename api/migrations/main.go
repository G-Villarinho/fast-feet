package main

import (
	"context"
	"log"
	"time"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/G-Villarinho/fast-feet-api/models"
	"github.com/G-Villarinho/fast-feet-api/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := storage.NewPostgresStorage(ctx)
	if err != nil {
		log.Fatal("error to connect to database: ", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Recipient{},
	); err != nil {
		log.Fatal("error to migrate: ", err)
	}

	log.Println("Migration executed successfully")

	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
	users := []models.User{
		{
			BaseModel: models.BaseModel{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
			},
			FullName:     "Admin User",
			CPF:          "84908118060",
			Email:        "admin@fastfeet.com",
			PasswordHash: "$2y$10$QtkenSL2ECTKogeczO21t.cZgFjwQj2hxxFAf2WL.4oaJnkMEY9SG",
			Status:       models.ActiveStatus,
			Role:         models.Admin,
		},
		{
			BaseModel: models.BaseModel{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
			},
			CPF:          "90654203040",
			FullName:     "Delivery User",
			Email:        "delivery@fastfeet.com",
			PasswordHash: "$2y$10$QtkenSL2ECTKogeczO21t.cZgFjwQj2hxxFAf2WL.4oaJnkMEY9SG",
			Status:       models.ActiveStatus,
			Role:         models.DeliveryMan,
		},
		{
			BaseModel: models.BaseModel{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
			},
			CPF:          "23503028064",
			FullName:     "Delivery User 1",
			Email:        "delivery1@fastfeet.com",
			PasswordHash: "$2y$10$QtkenSL2ECTKogeczO21t.cZgFjwQj2hxxFAf2WL.4oaJnkMEY9SG",
			Status:       models.ActiveStatus,
			Role:         models.DeliveryMan,
		},
	}

	for _, user := range users {
		var existingUser models.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			if err := db.Create(&user).Error; err != nil {
				log.Printf("Error creating user %s: %v", user.Email, err)
			} else {
				log.Printf("User %s created successfully", user.Email)
			}
		} else {
			log.Printf("User %s already exists, skipping...", user.Email)
		}
	}
}
