package main

import (
	"MyGram/models"
	"MyGram/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var PORT = "localhost:8080"

	dsn := "host=localhost user=postgres password=****** dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	if err != nil {
		panic(err)
	}

	routers.SetupRouter(db).Run(PORT)
}
