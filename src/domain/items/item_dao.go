package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dung997bn/bookstore_items_api/src/client/elasticsearch"
	"github.com/dung997bn/bookstore_items_api/src/domain/queries"
	"github.com/dung997bn/bookstore_utils-go/resterrors"
)

const (
	indexItems  = "items"
	docTypeItem = "_doc"
)

//Save func
func (i *Item) Save() resterrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, docTypeItem, i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

//GetByID func
func (i *Item) GetByID() resterrors.RestErr {
	itemID := i.ID
	result, err := elasticsearch.Client.GetByID(indexItems, docTypeItem, i.ID)
	if err != nil {
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to get data from id: %s", i.ID), errors.New("database error"))
	}
	if !result.Found {
		return resterrors.NewNotFoundError(fmt.Sprintf("No items found with id :%s", i.ID))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to parse json from result", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return resterrors.NewInternalServerError("error when trying to parse json from result", errors.New("database error"))
	}
	i.ID = itemID
	return nil
}

//Search func
func (i *Item) Search(query queries.EsQuery) ([]Item, resterrors.RestErr) {

	resultQuery, err := elasticsearch.Client.Search(indexItems, docTypeItem, query.Build())
	if err != nil {
		return nil, resterrors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	result := make([]Item, 0)
	if resultQuery.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", resultQuery.Hits.TotalHits)

		for _, hit := range resultQuery.Hits.Hits {
			// Deserialize hit.Source into a Item (could also be just a map[string]interface{}).
			var i Item
			bytes, err := hit.Source.MarshalJSON()
			if err != nil {
				return nil, resterrors.NewInternalServerError("error when trying to MarshalJSON documents", errors.New("database error"))
			}
			if err := json.Unmarshal(bytes, &i); err != nil {
				return nil, resterrors.NewInternalServerError("error when trying to Deserialize documents", errors.New("database error"))
			}
			i.ID = hit.Id
			result = append(result, i)
		}
	} else {
		// No hits
		return nil, resterrors.NewNotFoundError("no items found matching given condition")
	}
	if len(result) == 0 {
		return nil, resterrors.NewNotFoundError("no items found matching given condition")
	}
	return result, nil
}

//Delete fun
func (i *Item) Delete() (string, resterrors.RestErr) {
	if err := i.GetByID(); err != nil {
		return "", resterrors.NewNotFoundError("no items found matching given condition")
	}
	result, err := elasticsearch.Client.Delete(indexItems, docTypeItem, i.ID)
	if err != nil {
		return "", resterrors.NewInternalServerError(fmt.Sprintf("error when trying to delete document with id: %s", i.ID), errors.New("database error"))
	}
	if result.Result != "deleted" {
		return "", resterrors.NewInternalServerError(fmt.Sprintf("error when trying to delete document with id: %s", i.ID), errors.New("database error"))
	}
	return result.Result, nil
}

//Update func
func (i *Item) Update(updateBody *Item) resterrors.RestErr {
	if err := i.GetByID(); err != nil {
		return resterrors.NewNotFoundError("no items found matching given condition")
	}
	body := MakeUpdateBody(i, updateBody)
	result, err := elasticsearch.Client.Update(indexItems, docTypeItem, i.ID, body)

	if err != nil {
		return resterrors.NewInternalServerError(fmt.Sprintf("error when trying to update document with id: %s", i.ID), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to parse json from result", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return resterrors.NewInternalServerError("error when trying to parse json from result", errors.New("database error"))
	}
	return nil
}
