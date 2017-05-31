package archive

import (
	"testing"
)

func TestCountSources(t *testing.T) {
	// TODO - some test isn't cleaning up after itself.
	if err := resetTestData(appDB, "sources"); err != nil {
		t.Fatal(err)
	}

	c, err := CountSources(appDB)
	if err != nil {
		t.Error(err.Error())
	}

	if c != 4 {
		t.Errorf("wrong number of sources, expected %d, got: %d", 4, c)
	}
}

func TestListSources(t *testing.T) {
	s, err := ListSources(appDB, 100, 0)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(s) != 4 {
		t.Errorf("wrong number of sources, expected %d, got: %d", 4, len(s))
	}
}

func TestCrawlingSources(t *testing.T) {
	s, err := CrawlingSources(appDB, 100, 0)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(s) != 3 {
		t.Errorf("wrong number of crawling sources, expected %d, got: %d", 3, len(s))
	}
}
