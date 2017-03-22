package archive

import (
	"fmt"
	"testing"
)

func TestCollectionStorage(t *testing.T) {
	defer resetTestData(appDB, "collections")

	c := &Collection{Title: "test collection"}
	if err := c.Save(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	c.Creator = "penelope"
	if err := c.Save(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	c2 := &Collection{Id: c.Id}
	if err := c2.Read(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	if !c2.Created.Equal(c.Created) {
		t.Errorf("created doesn't match: %s != %s", c2.Created.String(), c.Created.String())
	}

	if !c2.Updated.Equal(c.Updated) {
		t.Errorf("updated doesn't match: %s != %s", c2.Updated.String(), c.Updated.String())
	}

	if err := c.Delete(appDB); err != nil {
		t.Error(err.Error())
		return
	}
}
