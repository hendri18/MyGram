package main

import (
	"MyGram/models"
	"MyGram/routers"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	servicePort := os.Getenv("PORT")
	if servicePort == "" {
		servicePort = "8080"
	}

	pgHost := os.Getenv("PG_HOST")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgPort := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		pgHost,
		pgUser,
		pgPassword,
		pgDB,
		pgPort)

	// dsn := "host=localhost user=postgres password=****** dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	if err != nil {
		panic(err)
	}

	routers.SetupRouter(db).Run(servicePort)
}
