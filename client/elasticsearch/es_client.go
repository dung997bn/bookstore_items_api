package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/dung997bn/bookstore_utils-go/logger"

	"github.com/olivere/elastic"
)

var (
	//Client type
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	GetByID(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
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

func (e *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := e.client.Search(index).
		Index(index). // search in index "twitter"
		Query(query). // specify the query
		//Sort("user", true). // sort by "user" field, ascending

		//Pretty(true).       // pretty print request and response JSON
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search document in index %s", index), err)
		return nil, err
	}
	return result, nil
}
