package services

import (
	"github.com/dung997bn/bookstore_items_api/domain/items"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

var (
	//ItemsService declare
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *resterrors.RestErr)
	Get(string) (*items.Item, *resterrors.RestErr)
}

type itemService struct{}

func (i *itemService) Create(item items.Item) (*items.Item, *resterrors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *itemService) Get(string) (*items.Item, *resterrors.RestErr) {
	return nil, nil
}
