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

	boostrap(client)

	router := fiber.New()
	server := api.ServerInterfaceImpl{
		Client:   client,
		Validate: validator.New(),
	}

	api.RegisterHandlers(router, &server)

	router.Listen(":3000")
}

func boostrap(client *ent.Client) {
	ctx := context.Background()

	user1, err := client.User.Create().
		SetFirstName("Ben").
		SetLastName("Tom").Save(ctx)

	if err != nil {
		log.Println("failed creating car: %w", err)
		return
	}

	log.Println("user created : ", user1)

	user2, err := client.User.Create().
		SetFirstName("Chris").
		SetLastName("Tom").Save(ctx)

	if err != nil {
		log.Println("failed creating car: %w", err)
		return
	}

	log.Println("user created : ", user2)

	org, err := client.Organisation.Create().
		SetName("Bunnings").
		AddUsers(user1, user2).Save(ctx)

	if err != nil {
		log.Println("failed creating org: %w", err)
		return
	}

	log.Println("organisation created : ", org)

}
