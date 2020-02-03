package firestore

import (
	"google.golang.org/api/iterator"
)

func (f *firestoreClient) ListenQueryAsync(ctx context.Context, collection, path, operation string, value interface{}, Listener FirebaseQueryUpdate) error {
	query := f.firestoreClient.Collection(collection)

	iter := query.Where(path, operation, value).Snapshots(ctx)
	go func(iter *firestore.QuerySnapshotIterator, Listener FirebaseQueryUpdate) {
		for {
			q, err := iterator.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				Listener(nil, err)
			}
			Listener(q, nil)
		}
	}(q, Listener)
	return nil
}
