package search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"platzi-cqrs/models"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearchRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticSearchRepository, error) {
	log.Println("URL para elastic")
	log.Println(url)
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{url},
	})

	if err != nil {
		return nil, err
	}

	return &ElasticSearchRepository{client: client}, nil
}

func (e *ElasticSearchRepository) Close() {
	//
}

func (e *ElasticSearchRepository) IndexFeed(ctx context.Context, feed models.Feed) error {
	body, _ := json.Marshal(feed)
	_, err := e.client.Index(
		"feeds",
		bytes.NewReader(body),
		e.client.Index.WithDocumentID(feed.ID),
		e.client.Index.WithContext(ctx),
		e.client.Index.WithRefresh("wait_for"),
	)

	return err
}

func (e *ElasticSearchRepository) SearchFeed(ctx context.Context, query string) (results []models.Feed, err error) {
	var buf bytes.Buffer

	log.Println("aqui en el query")
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"title", "description"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}

	log.Println("error del enconder")
	if err = json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, err
	}

	log.Println("muhcos contextos")
	res, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex("feeds"),
		e.client.Search.WithBody(&buf),
		e.client.Search.WithTrackTotalHits(true),
	)

	log.Println("errores de los contextos")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			results = nil
		}
	}()

	log.Println("error alla")
	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var eRes map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&eRes); err != nil {
		return nil, err
	}

	var feeds []models.Feed

	log.Println("leyendo todos los hits")
	for _, hit := range eRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
		feed := models.Feed{}
		source := hit.(map[string]interface{})["_source"]
		marshal, err := json.Marshal(source)

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(marshal, &feed); err == nil {
			feeds = append(feeds, feed)
		}
	}

	return feeds, nil
}
