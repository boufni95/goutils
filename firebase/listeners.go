package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (f *firebaseClient) ListenQueryAsync(ctx context.Context, collection, path, operation string, value interface{}, Listener QueryListener) error {
	query := f.firestoreClient.Collection(collection)

	iter := query.Where(path, operation, value).Snapshots(ctx)
	go func(iter *firestore.QuerySnapshotIterator, Listener QueryListener) {
		for {
			q, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				Listener(nil, err)
			}
			up := QueryUpdate{
				Update: q,
			}
			Listener(&up, nil)
		}
	}(iter, Listener)
	return nil
}
