package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/dung997bn/bookstore_items_api/logger"
	"github.com/olivere/elastic"
)

var (
	//Client type
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

//Init func
func Init() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetGzip(true),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		// elastic.SetHeaders(http.Header{
		// 	"X-Caller-Id": []string{"..."},
		// }),
	)

	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (e *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := e.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (e *esClient) setClient(client *elastic.Client) {
	e.client = client
}