package esclient

// es client: https://olivere.github.io/elastic/
import (
	"context"
	"fmt"
	"time"

	LOG "github.com/cihub/seelog"
	"gopkg.in/olivere/elastic.v5" // import elastic
)

type EsClient struct {
	client      *elastic.Client
	indexPrefix string
	initSuccess bool
	// bulkRequest *elastic.BulkService
}

const DefaultRecommendIndexPrefix = "recommend_videos_"

func NewEsClient(esAddr string, indexPrefix string) (*EsClient, error) {
	c := &EsClient{
		client:      nil,
		indexPrefix: indexPrefix,
		initSuccess: false,
	}

	client, err := elastic.NewClient(elastic.SetURL(esAddr))
	if err == nil {
		c.client = client
		c.initSuccess = true
		// c.bulkRequest = client.Bulk()
	}

	info, code, err := client.Ping(esAddr).Do(context.Background())
	if err != nil {
		// Handle error
		LOG.Errorf("Ping server failed")
	} else {
		LOG.Infof("ES Server info, info=%s, code=code", info, code)
	}
	return c, err
}

func (e *EsClient) IndexName() string {
	return e.IndexNameDelay(0)
}

func (e *EsClient) IndexNameDelay(days int) string {
	current := time.Now().AddDate(0, 0, days)
	date := current.Format("02_01_2006")
	index := fmt.Sprintf("%s_%s", e.indexPrefix, date)
	return index
}
