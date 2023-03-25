package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	if projectID == "" {
		projectID = "my-project-id"
	}

	fmt.Println("Project ID:", projectID)

	// Use emulator host if set, otherwise use production Firestore
	emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST")
	if emulatorHost == "" {
		emulatorHost = "127.0.0.1:8080"
	}

	fmt.Println("Emulator host:", emulatorHost)

	var client *firestore.Client
	var err error

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)

	client, err = firestore.NewClient(ctx, projectID, option.WithoutAuthentication(), option.WithEndpoint(emulatorHost))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	// Add a document
	_, err = client.Collection("users").Doc("user1").Set(ctx, map[string]interface{}{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("Failed to add document: %v", err)
	}

	// Read the document
	doc, err := client.Collection("users").Doc("user1").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get document: %v", err)
	}

	fmt.Println("Document data:", doc.Data())
}
