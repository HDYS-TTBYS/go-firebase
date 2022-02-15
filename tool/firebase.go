package tool

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

func FirebaseInit() (*firebase.App, error) {
	firebaseCredential := os.Getenv("FIREBASE_CREDENTIAL")
	decFirebaseCredential, err := base64.StdEncoding.DecodeString(firebaseCredential)
	if err != nil {
		log.Fatal(err)
	}
	opt := option.WithCredentialsFile(string(decFirebaseCredential))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing furebase app: %v", err)
	}
	return app, nil
}
