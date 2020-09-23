package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/dung997bn/bookstore_utils-go/logger"

	"gopkg.in/olivere/elastic.v6"
)

var (
	//Client type
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	GetByID(string, string, string) (*elastic.GetResult, error)
	Search(string, string, elastic.Query) (*elastic.SearchResult, error)
	Delete(string, string, string) (*elastic.DeleteResponse, error)
	Update(string, string, string, map[string]interface{}) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

//Init func
func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetGzip(true),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)

	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (e *esClient) setClient(client *elastic.Client) {
	e.client = client
}

func (e *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := e.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (e *esClient) GetByID(index string, docType string, ID string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := e.client.Get().
		Index(index).
		Id(ID).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", ID), err)
		return nil, err
	}

	return result, nil
}

func (e *esClient) Search(index string, docType string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := e.client.Search(index).
		Index(index). // search in index "twitter"
		Query(query). // specify the query
		Type(docType).
		//Sort("user", true). // sort by "user" field, ascending

		//Pretty(true).       // pretty print request and response JSON
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (e *esClient) Delete(index string, docType string, itemID string) (*elastic.DeleteResponse, error) {
	ctx := context.Background()
	result, err := e.client.Delete().
		Index(index).
		Type(docType).
		Id(itemID).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to delete docoment with id %s", itemID), err)
		return nil, err
	}
	return result, err
}

func (e *esClient) Update(index string, docType string, itemID string, updateBody map[string]interface{}) (*elastic.GetResult, error) {
	ctx := context.Background()
	update, err := e.client.Update().Index(index).Type(docType).Id(itemID).
		Doc(updateBody).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to update docoment with id %s", itemID), err)
	}
	if update.Result != "updated" {
		logger.Error(fmt.Sprintf("error when trying to update docoment with id %s", itemID), err)
	}
	return e.GetByID(index, docType, itemID)
}
