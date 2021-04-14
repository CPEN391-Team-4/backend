package notification

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"

	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// Send: Send a push notification to tok
func Send(tok string, title string, body string, credentialsFile string) (string, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", fmt.Errorf("error initializing app: %v", err)
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Messaging client: %v", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: tok,
	}

	// Send a message to the device corresponding to the provided
	// registration token.

	return client.Send(ctx, message)
}