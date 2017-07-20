package archive

import (
	"fmt"
	"github.com/datatogether/sql_datastore"
	"github.com/ipfs/go-datastore"
	"testing"
)

// confirm CollectionItem conforms to sql datastore mode
var _ = sql_datastore.Model(&CollectionItem{})

func TestCollectionItemStorage(t *testing.T) {
	store := datastore.NewMapDatastore()

	c := &Collection{Title: "test collection"}
	if err := c.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	u := Url{Url: "http://example.url.test.com"}

	item := &CollectionItem{collectionId: c.Id, Url: u, Index: 0, Description: "Description"}
	if err := item.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	item2 := &CollectionItem{collectionId: c.Id, Url: u}
	if err := item2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if err := CompareCollectionItems(item, item2); err != nil {
		t.Error(err.Error())
		return
	}

	item.Description = "Updated Description"
	if err := item.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	item2 = &CollectionItem{collectionId: c.Id, Url: u}
	if err := item2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if err := CompareCollectionItems(item, item2); err != nil {
		t.Error(err.Error())
		return
	}

	if err := item.Delete(store); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestCollectionItemSQLStorage(t *testing.T) {
	// defer resetTestData(appDB, "collections")

	// store := sql_datastore.NewDatastore(appDB)
	// if err := store.Register(&Collection{}); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	// c := &Collection{Title: "test collection"}
	// if err := c.Save(store); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	// c.Creator = "penelope"
	// if err := c.Save(store); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	// c2 := &Collection{Id: c.Id}
	// if err := c2.Read(store); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	// if err := CompareCollections(c, c2); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }

	// if err := c.Delete(store); err != nil {
	// 	t.Error(err.Error())
	// 	return
	// }
}

func CompareCollectionItems(a, b *CollectionItem) error {
	if err := CompareUrls(&a.Url, &b.Url); err != nil {
		return fmt.Errorf("url mismatch: %s", err.Error())
	}
	if a.collectionId != b.collectionId {
		return fmt.Errorf("collectionId mismatch: %d != %d", a.Index, b.Index)
	}
	if a.Index != b.Index {
		return fmt.Errorf("Index mismatch: %d != %d", a.Index, b.Index)
	}
	if a.Description != b.Description {
		return fmt.Errorf("Description mistmatch: %s != %s", a.Description, b.Description)
	}
	return nil
}
