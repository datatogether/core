package archive

import (
	"database/sql"
	"fmt"
	"github.com/datatogether/sql_datastore"
	"github.com/datatogether/sqlutil"
	"github.com/ipfs/go-datastore"
)

type CollectionItem struct {
	// need a reference to the collection Id to be set to distinguish
	// this item's membership in this particular list
	collectionId string
	// Collection Items are Url's at heart
	Url
	// this item's natural index in the colleciton
	Index int
	// unique description of this item
	Description string
}

func (c CollectionItem) DatastoreType() string {
	return "CollectionItem"
}

func (c CollectionItem) GetId() string {
	return c.Id
}

func (c CollectionItem) Key() datastore.Key {
	return datastore.NewKey(fmt.Sprintf("%s:%s/%s", c.DatastoreType(), c.collectionId, c.GetId()))
}

// Read collection from db
func (c *CollectionItem) Read(store datastore.Datastore) error {
	ci, err := store.Get(c.Key())
	if err != nil {
		return err
	}

	got, ok := ci.(*CollectionItem)
	if !ok {
		return ErrInvalidResponse
	}
	*c = *got
	return nil
}

// Save a collection
func (c *CollectionItem) Save(store datastore.Datastore) (err error) {
	// var exists bool

	// if c.Id != "" {
	// 	exists, err = store.Has(c.Key())
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// if !exists {
	// 	c.Id = uuid.New()
	// 	c.Created = time.Now().Round(time.Second)
	// 	c.Updated = c.Created
	// } else {
	// 	c.Updated = time.Now().Round(time.Second)
	// }

	return store.Put(c.Key(), c)
}

// Delete a collection, should only do for erronious additions
func (c *CollectionItem) Delete(store datastore.Datastore) error {
	return store.Delete(c.Key())
}

func (c *CollectionItem) NewSQLModel(key datastore.Key) sql_datastore.Model {
	l := key.List()
	if len(l) == 2 {
		return &CollectionItem{
			collectionId: l[0],
			Url:          Url{Id: l[1]},
		}
	}
	return &CollectionItem{}
}

func (c CollectionItem) SQLQuery(cmd sql_datastore.Cmd) string {
	switch cmd {
	case sql_datastore.CmdCreateTable:
		return qCollectionItemCreateTable
	case sql_datastore.CmdExistsOne:
		return qCollectionItemExists
	case sql_datastore.CmdSelectOne:
		return qCollectionItemById
	case sql_datastore.CmdInsertOne:
		return qCollectionItemInsert
	case sql_datastore.CmdUpdateOne:
		return qCollectionItemUpdate
	case sql_datastore.CmdDeleteOne:
		return qCollectionItemDelete
	case sql_datastore.CmdList:
		return qCollectionItems
	default:
		return ""
	}
}

func (c *CollectionItem) SQLParams(cmd sql_datastore.Cmd) []interface{} {
	switch cmd {
	case sql_datastore.CmdSelectOne, sql_datastore.CmdExistsOne, sql_datastore.CmdDeleteOne:
		return []interface{}{c.collectionId, c.Url.Id}
	case sql_datastore.CmdList:
		return nil
	default:
		return []interface{}{
			c.collectionId,
			c.Url.Id,
			c.Index,
			c.Description,
		}
	}
}

// UnmarshalSQL reads an sql response into the collection receiver
// it expects the request to have used collectionCols() for selection
func (c *CollectionItem) UnmarshalSQL(row sqlutil.Scannable) (err error) {
	var (
		collectionId, urlId, description string
		index                            int
	)

	if err := row.Scan(&collectionId, &urlId, &index, &description); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	*c = CollectionItem{
		collectionId: collectionId,
		Url:          Url{Id: urlId},
		Index:        index,
		Description:  description,
	}

	return nil
}
