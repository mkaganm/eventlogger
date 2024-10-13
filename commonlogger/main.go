package main

import (
	"commonlogger/internal/clients/mongo"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Create a new MongoDB client and connect to the 'testdb' database
	client, err := mongo.NewClient("testdb")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		// Close the MongoDB connection
		if err := client.Close(); err != nil {
			log.Fatalf("Error closing MongoDB connection: %v", err)
		}
	}()

	// Insert a new document into the 'users' collection
	doc := bson.M{
		"_id":  "12345",
		"name": "John Doe",
		"age":  30,
	}
	if err := client.InsertDocument("users", doc); err != nil {
		log.Fatalf("Error inserting document: %v", err)
	}
	fmt.Println("Document successfully inserted!")

	// Retrieve the inserted document by its ID
	retrievedDoc, err := client.GetDocument("users", "12345")
	if err != nil {
		log.Fatalf("Error retrieving document: %v", err)
	}
	fmt.Printf("Retrieved document: %v\n", retrievedDoc)

	// Update the document by its ID
	update := bson.M{
		"age": 31,
	}
	if err := client.UpdateDocument("users", "12345", update); err != nil {
		log.Fatalf("Error updating document: %v", err)
	}
	fmt.Println("Document successfully updated!")

	// Delete the document by its ID
	if err := client.DeleteDocument("users", "12345"); err != nil {
		log.Fatalf("Error deleting document: %v", err)
	}
	fmt.Println("Document successfully deleted!")
}
