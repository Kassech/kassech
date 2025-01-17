package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var FirebaseClient = (*messaging.Client)(nil)

func InitFirebaseClient() {
	opt := option.WithCredentialsFile("./kassechtransportaion-firebase-adminsdk-spm72-9c4f203fc9.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FirebaseClient, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
}
