package firebase

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func NewFirebaseClient(firestore, auth, messaging bool) Client {
	fc := firebaseClient{
		firestore: firestore,
		auth:      auth,
		messaging: messaging,
	}
	return &fc
}

func (f *firebaseClient) Connect(ctx context.Context, jsonAdminPath string) error {
	opt := option.WithCredentialsFile(jsonAdminPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	isInit := false
	if err != nil {
		return err
	}
	if f.auth {
		f.authClient, err = app.Auth(ctx)
		if err != nil {
			return err
		}
		isInit = true
	}
	if f.firestore {
		f.firestoreClient, err = app.Firestore(ctx)
		if err != nil {
			return err
		}
		isInit = true
	}
	if f.messaging {
		f.messagingClient, err = app.Messaging(ctx)
		if err != nil {
			return err
		}
		isInit = true
	}
	if !isInit {
		return errors.New("At least one service should be initialized")
	}

	return nil

}

func (f *firebaseClient) GetFirestoreClient() *firestore.Client {
	return f.firestoreClient
}
func (f *firebaseClient) GetAuthClient() *auth.Client {
	return f.authClient
}
func (f *firebaseClient) GetMessagingClient() *messaging.Client {
	return f.messagingClient
}
