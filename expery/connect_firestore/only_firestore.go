/*
	Environment Variable:
	The most common method is to set the GOOGLE_APPLICATION_CREDENTIALS environment variable to the path of your downloaded JSON key file. This allows the SDK to automatically discover your credentials.
	    Linux/macOS: export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/key.json"
	    Windows (PowerShell): $env:GOOGLE_APPLICATION_CREDENTIALS="C:\path\to\your\key.json"
	    Windows (Command Prompt): set GOOGLE_APPLICATION_CREDENTIALS=C:\path\to\your\key.json
	Explicit Credentials:
	Alternatively, you can provide the credentials directly in your code using the option.WithCredentialsFile() function.

Key Considerations:

	Project ID: Replace "your-project-id" with your actual Firebase project ID.
	Security: Keep your service account key file secure. Avoid committing it to version control.
	Environment: The GOOGLE_APPLICATION_CREDENTIALS environment variable is ideal for local development and testing. For production environments, consider using other secure methods for providing credentials (e.g., secrets management).

Note:
The Go SDK does not directly support user authentication using ID tokens, you would need to use the Firestore REST API for that.
*/
package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	// Option 1: Using GOOGLE_APPLICATION_CREDENTIALS environment variable
	client, err := firestore.NewClient(ctx, "your-project-id")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Option 2: Using explicit credentials file
	// credsFile := "/path/to/your/key.json"
	// options := option.WithCredentialsFile(credsFile)
	// client, err := firestore.NewClient(ctx, "your-project-id", options)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// Now you can use the client to interact with Firestore
	// Example:
	// _, _, err = client.Collection("your-collection").Add(ctx, map[string]interface{}{
	// 	"field": "value",
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to add document: %v", err)
	// }

	log.Println("Firestore client initialized successfully")
}
