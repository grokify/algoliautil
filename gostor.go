package algoliautil

import (
	"errors"
	"fmt"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

const noAlgoliaIndexSetError = "no Algolia index set"

type itemKeyValue struct {
	Key             string    `json:"key"`
	Value           string    `json:"value"`
	CreationTime    time.Time `json:"createTime"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}

// GostorClient is a client that satisfies the grokify/gostor.Client
// interface.
type GostorClient struct {
	client    *search.Client
	index     *search.Index
	indexName string
}

func NewGostorClient(client *search.Client, indexName string) GostorClient {
	return GostorClient{
		client:    client,
		indexName: indexName,
		index:     client.InitIndex(indexName)}
}

func (c GostorClient) SetString(key, val string) error {
	if c.index == nil {
		return errors.New(noAlgoliaIndexSetError)
	}
	_, err := c.index.SaveObject(itemKeyValue{
		Key:             key,
		Value:           val,
		CreationTime:    time.Now().UTC(),
		LastUpdatedTime: time.Now().UTC()})
	return err
}

func (c GostorClient) GetString(key string) (string, error) {
	if c.index == nil {
		return "", errors.New(noAlgoliaIndexSetError)
	}

	queryRes, err := c.index.Search(key, nil)
	if err != nil {
		return "", err
	}

	var items []itemKeyValue
	err = queryRes.UnmarshalHits(&items)
	if err != nil {
		return "", err
	}
	if len(items) != 1 {
		return "", fmt.Errorf("search returned !1 items [%d]", len(items))
	}
	return items[0].Value, nil
}

func (c GostorClient) MustGetString(key string) string {
	val, err := c.GetString(key)
	if err != nil {
		return ""
	}
	return val
}
