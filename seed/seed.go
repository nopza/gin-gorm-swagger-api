package seed

import (
	"gin-gorm-swagger-api/database"
	"gin-gorm-swagger-api/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	if database.DB == nil {
		log.Fatal("Database not initialized. Call ConnectDatabase() first.")
	}

	users := []models.User{
		{Name: "Admin", Email: "admin@example.com", Password: "admin123"},
		{Name: "Alice", Email: "alice@example.com", Password: "12345"},
	}

	for _, u := range users {
		var count int64
		database.DB.Model(&models.User{}).Where("email = ?", u.Email).Count(&count)

		if count == 0 {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			u.Password = string(hashedPassword)
			database.DB.Create(&u)
			log.Println("Seeded user:", u.Email)
		} else {
			log.Println("User already exists, skipping:", u.Email)
		}
	}
}
