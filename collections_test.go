package archive

import (
	"testing"
)

func TestListCollections(t *testing.T) {
	collections, err := ListCollections(appDB, 20, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(collections) != 1 {
		t.Errorf("collections length mismatch")
	}
}
