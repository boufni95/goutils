package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
)

type Client interface {
	GetFirestoreClient() *firestore.Client
	GetAuthClient() *auth.Client
	GetMessagingClient() *messaging.Client
	FirestoreGet(
		ctx context.Context,
		collection string,
		docID string,
	) (map[string]interface{}, error)
	FirestoreSet(
		ctx context.Context,
		collection,
		docID string,
		data map[string]interface{},
		merge bool,
	) error
	SendMessageTopic(
		ctx context.Context,
		topic string,
		data map[string]string,
	) error
	VerifyToken(
		ctx context.Context,
		token string,
	) (string, error)
	ListenQueryAsync(
		ctx context.Context,
		collection string,
		path string,
		operation string,
		value interface{},
		Listener FirebaseQueryUpdate,
	) error
}

type firebaseClient struct {
	firestore       bool
	auth            bool
	messaging       bool
	authClient      *auth.Client
	firestoreClient *firestore.Client
	messagingClient *messaging.Client
}

type FirebaseQueryListener func(update *FirebaseQueryUpdate, err error)

type FirebaseQueryUpdate struct {
	Update *firestore.QuerySnapshot
}
