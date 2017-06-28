package archive

import (
	"github.com/datatogether/sql_datastore"
	"github.com/ipfs/go-datastore"
	"testing"
)

func TestCollectionStorage(t *testing.T) {
	store := datastore.NewMapDatastore()

	c := &Collection{Title: "test collection"}
	if err := c.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	c.Creator = "penelope"
	if err := c.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	c2 := &Collection{Id: c.Id}
	if err := c2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if !c2.Created.Equal(c.Created) {
		t.Errorf("created doesn't match: %s != %s", c2.Created.String(), c.Created.String())
	}

	if !c2.Updated.Equal(c.Updated) {
		t.Errorf("updated doesn't match: %s != %s", c2.Updated.String(), c.Updated.String())
	}

	if err := c.Delete(store); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestCollectionSQLStorage(t *testing.T) {
	defer resetTestData(appDB, "collections")

	store := sql_datastore.NewDatastore(appDB)
	if err := store.Register(&Collection{}); err != nil {
		t.Error(err.Error())
		return
	}

	c := &Collection{Title: "test collection"}
	if err := c.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	c.Creator = "penelope"
	if err := c.Save(store); err != nil {
		t.Error(err.Error())
		return
	}

	c2 := &Collection{Id: c.Id}
	if err := c2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if !c2.Created.Equal(c.Created) {
		t.Errorf("created doesn't match: %s != %s", c2.Created.String(), c.Created.String())
	}

	if !c2.Updated.Equal(c.Updated) {
		t.Errorf("updated doesn't match: %s != %s", c2.Updated.String(), c.Updated.String())
	}

	if err := c.Delete(store); err != nil {
		t.Error(err.Error())
		return
	}
}
