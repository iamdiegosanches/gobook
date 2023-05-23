package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file", zap.Error(err))
	}
	return os.Getenv("MONGOURI")
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		logger.Fatal("error creating MongoDB client", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal("error connecting to MongoDB", zap.Error(err))
	}

	// ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("error pinging MongoDB", zap.Error(err))
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

// Query all books from the MongoDB collection.
func handleGetBooks(c *gin.Context, client *mongo.Client) ([]Book, error) {
	collection := GetCollection(client, "books")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var books []Book
	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, nil
}

func handlePostBooks(c *gin.Context, client *mongo.Client) {
	var newBook Book

	// call bind with validation
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new book with a generated ID
	book := NewBook(newBook.Title, newBook.Author, newBook.PublicationDate, newBook.Publisher)

	// Insert the new book into the MongoDB collection
	collection := GetCollection(client, "books")
	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to insert book"})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}

func handleGetById(c *gin.Context, client *mongo.Client) {
	uuidParam := c.Param("uuid")
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid uuid"})
		return
	}

	// Query the book by its ID from the MongoDB collection
	collection := GetCollection(client, "books")
	filter := bson.D{{Key: "id", Value: uuid}}
	var book Book
	err = collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to query book"})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func handleDeleteBook(c *gin.Context, client *mongo.Client) {
	uuidParam := c.Param("uuid")
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid uuid"})
		return
	}

	collection := GetCollection(client, "books")
	filter := bson.D{{Key: "id", Value: uuid}}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete book"})
		return
	}
	if deleteResult.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
}
