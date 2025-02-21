package database

import (
	"context"
	"fmt"
	"log"
	"os"
	// "time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	DB = client.Database(os.Getenv("DB_NAME"))
	fmt.Println("✅ Connected to MongoDB")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
