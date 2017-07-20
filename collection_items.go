package archive

import (
	"github.com/datatogether/sql_datastore"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// ItemCount gets the number of items in the list
func (c *Collection) ItemCount(store datastore.Datastore) (count int, err error) {
	if sqls, ok := store.(*sql_datastore.Datastore); ok {
		row := sqls.DB.QueryRow(qCollectionLength, c.Id)
		err = row.Scan(&count)
		return
	}

	// TODO - untested code :/
	res, err := store.Query(query.Query{
		Prefix:   c.Key().String(),
		KeysOnly: true,
	})
	if err != nil {
		return 0, err
	}

	for r := range res.Next() {
		if r.Error != nil {
			return 0, err
		}
		if _, ok := r.Value.(*CollectionItem); ok {
			count++
		}
	}

	return
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
		Prefix: c.Key().String(),
		Filters: []query.Filter{
			sql_datastore.FilterKeyTypeEq(CollectionItem{}.DatastoreType()),
		},
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
			return nil, r.Error
		}

		c, ok := r.Value.(*CollectionItem)
		if !ok {
			return nil, ErrInvalidResponse
		}

		items[i] = c
		i++
	}

	// fmt.Println(items)

	return items[:i], nil
}
