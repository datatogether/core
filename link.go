package archive

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/archivers-space/sqlutil"
	"github.com/ipfs/go-datastore"
	"github.com/multiformats/go-multihash"

	"time"
)

// A link represents an <a> tag in an html document src who's href
// attribute points to the url that resolves to dst.
// both src & dst must be stored as urls
type Link struct {
	// Calculated Hash for fixed ID purposes
	Hash string
	// created timestamp rounded to seconds in UTC
	Created time.Time `json:"created"`
	// updated timestamp rounded to seconds in UTC
	Updated time.Time `json:"updated"`
	// origin url of the linking document
	Src *Url `json:"src"`
	// absolute url of the <a> href property
	Dst *Url `json:"dst"`
}

func (l *Link) DatastoreType() string {
	return "Link"
}

func (l *Link) GetId() string {
	if l.Hash == "" {
		l.calcHash()
	}
	return l.Hash
}

func (l *Link) Key() datastore.Key {
	return datastore.NewKey(fmt.Sprintf("%s:%s", l.DatastoreType(), l.GetId()))
}

func (l *Link) Read(store datastore.Datastore) error {
	if l.Src == nil || l.Dst == nil {
		return ErrNotFound
	}

	li, err := store.Get(l.Key())
	if err != nil {
		return err
	}

	got, ok := li.(*Link)
	if !ok {
		return ErrInvalidResponse
	}

	*l = *got
	return nil
}

func (l *Link) Insert(store datastore.Datastore) error {
	l.Created = time.Now().In(time.UTC).Round(time.Second)
	l.Updated = l.Created
	return store.Put(l.Key(), l)
}

func (l *Link) Update(store datastore.Datastore) error {
	l.Updated = time.Now().Round(time.Second)
	return store.Put(l.Key(), l)
}

func (l *Link) Delete(store datastore.Datastore) error {
	return store.Delete(l.Key())
}

func (l *Link) calcHash() {
	h := sha256.New()
	data, err := json.Marshal(struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	}{
		Src: l.Src.Url,
		Dst: l.Dst.Url,
	})
	if err != nil {
		return
	}

	h.Write(data)
	mhBuf, err := multihash.EncodeName(h.Sum(nil), "sha2-256")
	if err != nil {
		return
	}

	l.Hash = hex.EncodeToString(mhBuf)
}

func (l Link) NewSQLModel(id string) sqlutil.Model {
	return &Link{Hash: id}
}

func (l *Link) SQLQuery(cmd sqlutil.CmdType) string {
	switch cmd {
	case sqlutil.CmdCreateTable:
		return qLinkCreateTable
	case sqlutil.CmdExistsOne:
		return qLinkExists
	case sqlutil.CmdInsertOne:
		return qLinkInsert
	case sqlutil.CmdDeleteOne:
		return qLinkDelete
	case sqlutil.CmdUpdateOne:
		return qLinkUpdate
	default:
		return ""
	}
}

func (l *Link) SQLParams(cmd sqlutil.CmdType) []interface{} {
	// TODO remove the need for these
	if l.Src == nil {
		l.Src = &Url{}
	}
	if l.Dst == nil {
		l.Dst = &Url{}
	}

	switch cmd {
	case sqlutil.CmdSelectOne, sqlutil.CmdExistsOne, sqlutil.CmdDeleteOne:
		return []interface{}{
			l.Src.Url,
			l.Dst.Url,
		}
	default:
		return []interface{}{
			l.Created.In(time.UTC),
			l.Updated.In(time.UTC),
			l.Src.Url,
			l.Dst.Url,
		}
	}
}

func (l *Link) UnmarshalSQL(row sqlutil.Scannable) error {
	var (
		created, updated time.Time
		src, dst         string
	)

	if err := row.Scan(&created, &updated, &src, &dst); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	*l = Link{
		Created: created.In(time.UTC),
		Updated: updated.In(time.UTC),
		Src:     &Url{Url: src},
		Dst:     &Url{Url: dst},
	}

	return nil
}
