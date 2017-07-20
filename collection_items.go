package archive

import (
	"github.com/datatogether/sql_datastore"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// ItemCount gets the number of items in the list
// TODO
func (c *Collection) ItemCount(store datastore.Datastore) (int, error) {
	return 0, nil
}

func (c *Collection) SaveItems(store datastore.Datastore, items []*CollectionItem) error {
	for _, item := range items {
		item.collectionId = c.Id
		if err := item.Save(store); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection) DeleteItems(store datastore.Datastore, items []*CollectionItem) error {
	for _, item := range items {
		item.collectionId = c.Id
		if err := item.Delete(store); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection) ReadItems(store datastore.Datastore, orderby string, limit, offset int) (items []*CollectionItem, err error) {
	items = make([]*CollectionItem, limit)

	res, err := store.Query(query.Query{
		Limit:  limit,
		Offset: offset,
		Orders: []query.Order{
			query.OrderByValue{
				TypedOrder: sql_datastore.OrderBy(orderby),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	i := 0
	for r := range res.Next() {
		if r.Error != nil {
			return nil, err
		}

		c, ok := r.Value.(*CollectionItem)
		if !ok {
			return nil, ErrInvalidResponse
		}

		items[i] = c
		i++
	}

	return items[:i], nil
}
