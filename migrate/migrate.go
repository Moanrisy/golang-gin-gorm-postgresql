package main

import (
	"fmt"
	"log"

	"github.com/moanrisy/golang-gorm-postgres/initializers"
	"github.com/moanrisy/golang-gorm-postgres/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Post{})
	// initializers.DB.Exec("DELETE FROM users")
	fmt.Println("? Migration complete")
}
