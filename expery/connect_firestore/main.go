package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Path to your service account key file
	credFile := "path/to/your/serviceAccountKey.json"

	// Initialize Firebase app
	opt := option.WithCredentialsFile(credFile)
	conf := &firebase.Config{ProjectID: "your-project-id"} // Replace with your project ID
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Initialize Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore client: %v\n", err)
	}
	defer client.Close()

	log.Println("Successfully connected to Firestore!")

	// You can now perform Firestore operations using the 'client' variable
}
