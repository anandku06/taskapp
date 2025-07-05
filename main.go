package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json : "id"`
	Completed bool   `json : "completed"`
	Body      string `json : "body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello World!!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	mongoDB := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoDB)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!!")

	collection = client.Database("task").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", createTodos)
	// app.Patch("/api/todos/:id", updateTodos)
	// app.Delete("/api/todos/:id", deleteTodos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {

}
// func createTodos(c *fiber.Ctx) error {}
// func updateTodos(c *fiber.Ctx) error {}
// func deleteTodos(c *fiber.Ctx) error {}
