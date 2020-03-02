package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/messaging"
)

func (f *firebaseClient) FirestoreDocExists(ctx context.Context, collection string, docID string) (bool, error) {
	doc, err := f.firestoreClient.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return false, err
	}
	return doc.Exists(), nil
}

func (f *firebaseClient) FirestoreGet(ctx context.Context, collection string, docID string) (map[string]interface{}, error) {
	if !f.firestore {
		return nil, &UnexpectedUseFirestoreErr{}
	}
	doc, err := f.firestoreClient.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc.Data(), nil
}

func (f *firebaseClient) FirestoreGetType(ctx context.Context, collection string, docID string, dataTo interface{}) error {
	if !f.firestore {
		return &UnexpectedUseFirestoreErr{}
	}
	doc, err := f.firestoreClient.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(dataTo)
}

func (f *firebaseClient) FirestoreSet(ctx context.Context, collection, docID string, data map[string]interface{}, merge bool) error {
	if !f.firestore {
		return &UnexpectedUseFirestoreErr{}
	}
	var err error
	if merge {
		_, err = f.firestoreClient.Collection(collection).Doc(docID).Set(ctx, data, firestore.MergeAll)
	} else {
		_, err = f.firestoreClient.Collection(collection).Doc(docID).Set(ctx, data)
	}
	return err
}
func (f *firebaseClient) FirestoreSetType(ctx context.Context, collection, docID string, dataTo interface{}, merge bool) error {
	if !f.firestore {
		return &UnexpectedUseFirestoreErr{}
	}
	var err error
	if merge {
		_, err = f.firestoreClient.Collection(collection).Doc(docID).Set(ctx, dataTo, firestore.MergeAll)
	} else {
		_, err = f.firestoreClient.Collection(collection).Doc(docID).Set(ctx, dataTo)
	}
	return err
}

func (f *firebaseClient) SendMessageTopic(ctx context.Context, topic string, data map[string]string) error {
	if !f.messaging {
		return &UnexpectedUseMessagingErr{}
	}
	message := &messaging.Message{
		Data:  data,
		Topic: topic,
	}

	_, err := f.messagingClient.Send(ctx, message)
	return err
}

func (f *firebaseClient) VerifyToken(ctx context.Context, token string) (string, error) {
	if !f.auth {
		return "", &UnexpectedUseAuthErr{}
	}
	tokenAuth, err := f.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}
	uid, ok := tokenAuth.Claims["user_id"].(string)
	if !ok {
		return "", &NoUIDFoundErr{}
	}
	return uid, nil
}
