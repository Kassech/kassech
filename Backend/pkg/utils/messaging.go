package utils

import (
	"context"
	"fmt"
	"kassech/backend/pkg/service"
	"log"

	"firebase.google.com/go/v4/messaging"
)

// SendMessage sends a single message to a device using its registration token.
func SendMessage(registrationToken, title, body string) {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: registrationToken,
	}

	response, err := service.FirebaseClient.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message:", response)
}

// SendMulticastMessage sends a message to multiple devices using their registration tokens.
func SendMulticastMessage(registrationTokens []string, title, body string) {
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: registrationTokens,
	}

	response, err := service.FirebaseClient.SendEachForMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Successfully sent multicast message to %d devices: %d succeeded, %d failed\n",
		len(registrationTokens), response.SuccessCount, response.FailureCount)
}
