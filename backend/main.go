package main

import (
	"context"
	"log"

	"github.com/adiwenak/hrapp/api"
	"github.com/adiwenak/hrapp/ent"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := fiber.New()
	server := api.ServerInterfaceImpl{
		Client:   client,
		Validate: validator.New(),
	}

	api.RegisterHandlers(router, &server)

	router.Listen(":3000")
}
