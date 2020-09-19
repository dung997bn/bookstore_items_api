package items

import (
	"errors"

	"github.com/dung997bn/bookstore_items_api/client/elasticsearch"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

const (
	indexItems = "items"
)

//Save func
func (i *Item) Save() *resterrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}
