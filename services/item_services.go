package services

import (
	"github.com/dung997bn/bookstore_items_api/domain/items"
	"github.com/dung997bn/bookstore_items_api/domain/queries"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

var (
	//ItemsService declare
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, resterrors.RestErr)
	GetByID(string) (*items.Item, resterrors.RestErr)
	Search(queries.EsQuery) ([]items.Item, resterrors.RestErr)
}

type itemService struct{}

func (i *itemService) Create(item items.Item) (*items.Item, resterrors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *itemService) GetByID(ID string) (*items.Item, resterrors.RestErr) {
	item := items.Item{
		ID: ID,
	}
	if err := item.GetByID(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *itemService) Search(query queries.EsQuery) ([]items.Item, resterrors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
