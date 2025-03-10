package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"repartnerstask.com/m/internal/config"
	"repartnerstask.com/m/internal/domain"
	"repartnerstask.com/m/internal/handlers"
	"repartnerstask.com/m/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("%v. Exiting...", err)
	}

	fmt.Printf("Pack sizes loaded from config: %v\n", cfg.PackSizes)

	repo := repository.NewInMemoryStorage()

	application, err := domain.NewApplication(cfg.PackSizes, repo)
	if err != nil {
		log.Fatalf("error initializing application: %v", err)
	}

	server := handlers.NewServer(&application)

	r := gin.Default()
	// disabling CORS for simplicity
	// in real-life applications CORS should be configured strictly, to only allow known request origins
	corsConfig := newCorseConfig()
	r.Use(cors.New(corsConfig))
	r.POST("/orders", server.PostOrders)
	r.GET("/orders", server.GetOrders)

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "password",
	}))
	authorized.PUT("/packs", server.PutPacks)
	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func newCorseConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")
	return corsConfig
}
