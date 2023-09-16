package main

import (
	"context"
	"log"

	"github.com/adiwenak/hrapp/ent"
	"github.com/adiwenak/hrapp/handlers"
	"github.com/adiwenak/hrapp/server"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {

	client, err := ent.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := fiber.New()

	core := &handlers.HRCore{
		Client:   client,
		Validate: validator.New(),
	}

	coreStrictHandler := server.NewStrictHandler(core, nil)

	server.RegisterHandlers(router, coreStrictHandler)

	// server := rest_handlers.ServerInterfaceImpl{
	// 	Client:   client,
	// 	Validate: validator.New(),
	// }

	// api.RegisterHandlers(router, &server)

	router.Listen(":4000")
}
