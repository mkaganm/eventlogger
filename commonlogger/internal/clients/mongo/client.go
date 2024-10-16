package mongo

import (
	"commonlogger/internal/constants"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// Client struct: Structure to manage MongoDB connection and basic operations.
type Client struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewClient Creates a new MongoClient and connects to MongoDB.
func NewClient(databaseName string) (*Client, error) {
	// DSN connection string (username, password, host, port)
	dsn := os.Getenv(constants.MONGO_DSN_URL)

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(dsn)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to check if the connection is active
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB ping failed: %v", err)
	}

	// Get a handle for the database
	database := client.Database(databaseName)

	return &Client{
		Client:   client,
		Database: database,
	}, nil
}

// InsertDocument Inserts a new document into the specified collection.
func (m *Client) InsertDocument(collectionName string, doc interface{}) error {
	collection := m.Database.Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	return nil
}

// GetDocument Retrieves a document by ID.
func (m *Client) GetDocument(collectionName string, docID string) (bson.M, error) {
	collection := m.Database.Collection(collectionName)

	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve document: %v", err)
	}

	return result, nil
}

// UpdateDocument Updates a document with the specified ID.
func (m *Client) UpdateDocument(collectionName string, docID string, update bson.M) error {
	collection := m.Database.Collection(collectionName)

	// Update the document
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": docID}, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}

// DeleteDocument Deletes a document by ID.
func (m *Client) DeleteDocument(collectionName string, docID string) error {
	collection := m.Database.Collection(collectionName)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": docID})
	if err != nil {
		return fmt.Errorf("failed to delete document: %v", err)
	}
	return nil
}

// Close Closes the MongoDB connection.
func (m *Client) Close() error {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %v", err)
	}
	return nil
}
