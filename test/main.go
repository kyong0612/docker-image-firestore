package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {

	FIRESTORE_EMULATOR_HOST := os.Getenv("FIRESTORE_EMULATOR_HOST")
	fmt.Println("Emulator host:", FIRESTORE_EMULATOR_HOST)

	var client *firestore.Client
	var err error

	ctx := context.Background()

	client, err = firestore.NewClient(
		ctx,
		"my-project-id",
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	fmt.Println("firestore client created")

	// Add a document
	_, err = client.Collection("users").Doc("user1").Set(ctx, map[string]interface{}{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("Failed to add document: %v", err)
	}

	fmt.Println("collection set (user1)")

	// Read the document
	doc, err := client.Collection("users").Doc("user1").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get document: %v", err)
	}

	fmt.Println("Document data:", doc.Data())
}
