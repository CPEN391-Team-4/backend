package main
import (
    "fmt"
    "context"

    firebase "firebase.google.com/go"

    "google.golang.org/api/option"
    "firebase.google.com/go/messaging"
    "log"
    "os"
)
func sendNotification(tok string, title string, body string) {
    path := "/home/aniket/go/src/github.com/CPEN391-Team-4/backend/src/notification/key.json"

    opt := option.WithCredentialsFile(path)
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v", err)
    }

    ctx := context.Background()
    client, err := app.Messaging(ctx)
    if err != nil {
        log.Fatalf("error getting Messaging client: %v\n", err)
    }


    // This registration token comes from the client FCM SDKs.
    //registrationToken := "cRfvd9qkRp6zhkjI8rwwF8:APA91bHO19zsujdKPIkdBxaddI-YIoqzS-UwmibR7gtiVPNzbuhbD-FL15Dbh_jBCumRisq2Slxa24iv7-EhPKXRL4KEqMz2dT_RILaYhqajhyxE6nufaL46aWNAHepITkOwFdtGwt5o"

    // See documentation on defining a message payload.
    message := &messaging.Message{
        Notification: &messaging.Notification{
            Title: title,
            Body:  body,
        },
        Token: tok,
    }

    // Send a message to the device corresponding to the provided
    // registration token.
    response, err := client.Send(ctx, message)
    if err != nil {
        log.Fatalln(err)
    }
    // Response is a message ID string.
    fmt.Println("Successfully sent message:", response)
}

func main() {
    sendNotification(os.Args[1], os.Args[2], os.Args[3])
}
