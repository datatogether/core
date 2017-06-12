package archive

import (
	// "fmt"
	"github.com/archivers-space/sql_datastore"
	"github.com/ipfs/go-datastore"
	"testing"
)

func TestUrlStorage(t *testing.T) {
	store := datastore.NewMapDatastore()

	u := &Url{Url: "http://youtube.com"}
	if err := u.Insert(store); err != nil {
		t.Error(err.Error())
		return
	}

	u.ContentLength = 10
	if err := u.Update(store); err != nil {
		t.Error(err.Error())
		return
	}

	u2 := &Url{Id: u.Id}
	if err := u2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if !u2.Created.Equal(u.Created) {
		t.Errorf("created doesn't match: %s != %s", u2.Created.String(), u.Created.String())
	}

	if !u2.Updated.Equal(u.Updated) {
		t.Errorf("updated doesn't match: %s != %s", u2.Updated.String(), u.Updated.String())
	}

	if err := u.Delete(store); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestUrlSQLStorage(t *testing.T) {
	defer resetTestData(appDB, "urls", "links")

	store := sql_datastore.NewDatastore(appDB)
	if err := store.Register(&Url{}); err != nil {
		t.Error(err)
		return
	}

	u := &Url{Url: "http://youtube.com"}
	if err := u.Insert(store); err != nil {
		t.Error(err.Error())
		return
	}

	u.ContentLength = 10
	if err := u.Update(store); err != nil {
		t.Error(err.Error())
		return
	}

	u2 := &Url{Url: "http://youtube.com"}
	if err := u2.Read(store); err != nil {
		t.Error(err.Error())
		return
	}

	if !u2.Created.Equal(u.Created) {
		t.Errorf("created doesn't match: %s != %s", u2.Created.String(), u.Created.String())
	}

	if !u2.Updated.Equal(u.Updated) {
		t.Errorf("updated doesn't match: %s != %s", u2.Updated.String(), u.Updated.String())
	}

	if err := u.Delete(store); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestShouldEnqueue(t *testing.T) {
	defer resetTestData(appDB, "urls")

	store := sql_datastore.NewDatastore(appDB)
	if err := store.Register(&Url{}); err != nil {
		t.Error(err)
		return
	}

	epa := &Url{Url: "http://www.epa.gov"}
	if err := epa.Read(store); err != nil {
		t.Fatal(err.Error())
	}

	cases := []struct {
		url       *Url
		get, head bool
	}{
		// TODO - this test isn't working the func properly. Should enhance with DB interaction
		{&Url{Url: "https://youtube.com"}, true, true},
		{&Url{Url: "http://www.fda.gov"}, true, true},
		{&Url{Url: "http://epa.gov/new"}, true, true},
		{epa, true, true},
	}

	for _, c := range cases {
		u := c.url
		head := u.ShouldEnqueueHead()
		if head != c.head {
			t.Errorf("shouldEnqueueHead: %s error. expected %t, got %t", u.Url, c.head, head)
		}

		get := u.ShouldEnqueueGet()
		if get != c.get {
			t.Errorf("shouldEnqueueGet: %s expected %t, got %t", u.Url, c.get, get)
		}
	}
}

// func TestUrlGet(t *testing.T) {
// 	u := &Url{Url: "https://www.apple.com"}
// 	done := make(chan bool, 0)
// 	links, err := u.Get(appDB, func(err error) {
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		done <- true
// 	})

// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if len(links) == 0 {
// 		t.Error("didn't find any links?")
// 	}
// 	<-done
// }

func TestUrlSuspectedContentUrl(t *testing.T) {
	cases := []struct {
		url    *Url
		expect bool
	}{
		{&Url{
			Url:           "https://opendata.epa.gov/home.xhtml?view",
			Status:        200,
			ContentType:   "text/html;charset=UTF-8",
			ContentLength: 18327,
			ContentSniff:  "text/plain; charset=utf-8",
		}, false},
	}

	for i, c := range cases {
		if got := c.url.SuspectedContentUrl(); got != c.expect {
			t.Errorf("case %d fail: %t != %t", i, c.expect, got)
		}
	}
}
