package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"web_crawler_app/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewDatabase(connectionString, dbName string) (*Database, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping the database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to MongoDB!")

	// Get database
	database := client.Database(dbName)

	return &Database{
		client:   client,
		database: database,
	}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.client.Disconnect(ctx)
}

// SaveIndex saves the index to the database
func (db *Database) SaveIndex(index *models.InvertedIndex) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get collection
	collection := db.database.Collection("index")

	// Delete existing index
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete existing index: %v", err)
	}

	// Insert new index
	indexDoc := bson.M{
		"index":     index.Index,
		"doc_freq":  index.DocFreq,
		"doc_count": index.DocCount,
	}

	_, err = collection.InsertOne(ctx, indexDoc)
	if err != nil {
		return fmt.Errorf("failed to insert index: %v", err)
	}

	return nil
}

// LoadIndex loads the index from the database
func (db *Database) LoadIndex() (*models.InvertedIndex, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get collection
	collection := db.database.Collection("index")

	// Find index
	var result bson.M
	err := collection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.NewInvertedIndex(), nil
		}
		return nil, fmt.Errorf("failed to load index: %v", err)
	}

	// Convert to index
	index := models.NewInvertedIndex()

	// Extract index data
	if indexData, ok := result["index"].(bson.M); ok {
		for term, docs := range indexData {
			if docsMap, ok := docs.(bson.M); ok {
				index.Index[term] = make(map[string]int)
				for docURL, freq := range docsMap {
					if freqInt, ok := freq.(int32); ok {
						index.Index[term][docURL] = int(freqInt)
					}
				}
			}
		}
	}

	// Extract doc_freq data
	if docFreqData, ok := result["doc_freq"].(bson.M); ok {
		for term, freq := range docFreqData {
			if freqInt, ok := freq.(int32); ok {
				index.DocFreq[term] = int(freqInt)
			}
		}
	}

	// Extract doc_count
	if docCount, ok := result["doc_count"].(int32); ok {
		index.DocCount = int(docCount)
	}

	return index, nil
}

// SaveDocuments saves documents to the database
func (db *Database) SaveDocuments(documents []models.Document) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get collection
	collection := db.database.Collection("documents")

	// Convert documents to BSON
	docs := make([]interface{}, len(documents))
	for i, doc := range documents {
		docs[i] = bson.M{
			"url":      doc.URL,
			"title":    doc.Title,
			"content":  doc.Content,
			"keywords": doc.Keywords,
		}
	}

	// Insert documents
	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return fmt.Errorf("failed to insert documents: %v", err)
	}

	return nil
}

// GetDocumentByURL retrieves a document by URL
func (db *Database) GetDocumentByURL(url string) (models.Document, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get collection
	collection := db.database.Collection("documents")

	// Find document
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"url": url}).Decode(&result)
	if err != nil {
		return models.Document{}, fmt.Errorf("failed to find document: %v", err)
	}

	// Convert to document
	document := models.Document{
		URL:     result["url"].(string),
		Title:   result["title"].(string),
		Content: result["content"].(string),
	}

	// Extract keywords
	if keywordsArray, ok := result["keywords"].(bson.A); ok {
		keywords := make([]string, len(keywordsArray))
		for i, keyword := range keywordsArray {
			keywords[i] = keyword.(string)
		}
		document.Keywords = keywords
	}

	return document, nil
}
