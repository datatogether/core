package archive

import (
	"testing"
)

func TestListSources(t *testing.T) {
	s, err := ListSources(appDB, 100, 0)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(s) != 4 {
		t.Errorf("wrong number of sources, expected %d, got: %d", 3, len(s))
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
