package couchbaseclient

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"time"
)

// CouchbaseClient struct
type CouchbaseClient struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
}

// NewCouchbaseClient initializes a new Couchbase client
func NewCouchbaseClient(connectionString, username, password, bucketName string) (*CouchbaseClient, error) {
	// Connect to Couchbase cluster
	cluster, err := gocb.Connect(connectionString, gocb.ClusterOptions{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, fmt.Errorf("cluster connection failed: %v", err)
	}

	// Get bucket connection
	bucket := cluster.Bucket(bucketName)

	// Wait for the bucket to be ready
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, fmt.Errorf("bucket connection failed: %v", err)
	}

	return &CouchbaseClient{Cluster: cluster, Bucket: bucket}, nil
}

// InsertDocument creates a new document in Couchbase
func (c *CouchbaseClient) InsertDocument(collectionName, docID string, doc interface{}) error {
	collection := c.Bucket.DefaultCollection()
	_, err := collection.Insert(docID, doc, nil)
	if err != nil {
		return fmt.Errorf("document insert failed: %v", err)
	}
	return nil
}

// GetDocument retrieves a document by ID
func (c *CouchbaseClient) GetDocument(collectionName, docID string) (interface{}, error) {
	collection := c.Bucket.DefaultCollection()
	getResult, err := collection.Get(docID, nil)
	if err != nil {
		return nil, fmt.Errorf("document retrieval failed: %v", err)
	}

	var doc interface{}
	err = getResult.Content(&doc)
	if err != nil {
		return nil, fmt.Errorf("failed to decode document: %v", err)
	}

	return doc, nil
}

// UpdateDocument updates an existing document
func (c *CouchbaseClient) UpdateDocument(collectionName, docID string, doc interface{}) error {
	collection := c.Bucket.DefaultCollection()
	_, err := collection.Replace(docID, doc, nil)
	if err != nil {
		return fmt.Errorf("document update failed: %v", err)
	}
	return nil
}

// DeleteDocument deletes a document by ID
func (c *CouchbaseClient) DeleteDocument(collectionName, docID string) error {
	collection := c.Bucket.DefaultCollection()
	_, err := collection.Remove(docID, nil)
	if err != nil {
		return fmt.Errorf("document delete failed: %v", err)
	}
	return nil
}
