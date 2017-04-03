package archive

import (
	"testing"
)

func TestUncrawlableStorage(t *testing.T) {
	defer resetTestData(appDB, "uncrawlables")

	p := &Uncrawlable{Url: "http://www.epa.gov"}
	if err := p.Save(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	p.Comments = "new comment"
	if err := p.Save(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	p2 := &Uncrawlable{Url: "http://www.epa.gov"}
	if err := p2.Read(appDB); err != nil {
		t.Error(err.Error())
		return
	}

	if !p2.Created.Equal(p.Created) {
		t.Errorf("created doesn't match: %s != %s", p2.Created.String(), p.Created.String())
	}

	if !p2.Updated.Equal(p.Updated) {
		t.Errorf("updated doesn't match: %s != %s", p2.Updated.String(), p.Updated.String())
	}

	if err := p.Delete(appDB); err != nil {
		t.Error(err.Error())
		return
	}
}
