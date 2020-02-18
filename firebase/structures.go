package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
)

//Client - The client interface with all implemented requests and subs ~WIP
type Client interface {
	//Connect - init all the requires clients
	Connect(
		ctx context.Context,
		jsonAdminPath string,
	) error
	//GetFirestoreClient - return the firestore client
	GetFirestoreClient() *firestore.Client
	//GetAuthClient - return the auth client
	GetAuthClient() *auth.Client
	//GetMessagingClient - return the messaging client
	GetMessagingClient() *messaging.Client
	//FirestoreGet - get a document from collection and ID
	FirestoreGet(
		ctx context.Context,
		collection string,
		docID string,
	) (map[string]interface{}, error)
	//FirestoreGetType - get a document from collection and ID into a typed param
	FirestoreGetType(
		ctx context.Context,
		collection string,
		docID string,
		dataTo interface{},
	) error
	//FirestoreSet - Set a document given collection and ID
	FirestoreSet(
		ctx context.Context,
		collection,
		docID string,
		data map[string]interface{},
		merge bool,
	) error
	//FirestoreSet - Set a document given collection and ID
	FirestoreSetType(
		ctx context.Context,
		collection,
		docID string,
		dataTo interface{},
		merge bool,
	) error
	//SendMessageTopic - Send message on topic
	SendMessageTopic(
		ctx context.Context,
		topic string,
		data map[string]string,
	) error
	//VerifyToken - verify auth token
	VerifyToken(
		ctx context.Context,
		token string,
	) (string, error)
	//ListenQueryAsync - listen for query updates
	ListenQueryAsync(
		ctx context.Context,
		collection string,
		path string,
		operation string,
		value interface{},
		Listener QueryListener,
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

//QueryListener - listen for query updates rends a response or an error
type QueryListener func(update *QueryUpdate, err error)

//QueryUpdate -
type QueryUpdate struct {
	Update *firestore.QuerySnapshot
}
