package archive

import (
	"github.com/archivers-space/sql_datastore"
	"testing"
)

func TestListCollections(t *testing.T) {
	store := sql_datastore.NewDatastore(appDB)
	if err := store.Register(&Collection{}); err != nil {
		t.Error(err.Error())
		return
	}

	collections, err := ListCollections(store, 20, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(collections) != 1 {
		t.Errorf("collections length mismatch")
	}

	collections, err = ListCollections(store, 20, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(collections) != 0 {
		t.Errorf("collections length mismatch")
	}
}
