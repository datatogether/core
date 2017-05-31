package archive

import (
	"testing"
)

func TestCountPrimers(t *testing.T) {
	// TODO - some test isn't cleaning up after itself.
	if err := resetTestData(appDB, "primers"); err != nil {
		t.Fatal(err)
	}

	c, err := CountPrimers(appDB)
	if err != nil {
		t.Error(err.Error())
	}

	if c != 6 {
		t.Errorf("wrong number of primers, expected %d, got: %d", 6, c)
	}
}
