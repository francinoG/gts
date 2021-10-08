package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Username string
	Password string
	Address  string
	Port     int
	DBName   string
}

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func DB() *mongo.Client {
	connectionString := "mongodb://" + Env("DB_USERNAME") + ":" + Env("DB_PASSWORD") + "@" + Env("DB_HOST") + ":" + Env("DB_PORT") + "/" + Env("DB_NAME")
	clientOptions := options.Client().ApplyURI(connectionString) // Connect to //MongoDB
	// clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017/gts") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
