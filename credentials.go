package algoliautil

import (
	"encoding/json"
	"errors"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

type Credentials struct {
	ApplicationId    string `json:"applicationId"`
	SearchOnlyApiKey string `json:"searchOnlyApiKey"`
	AdminApiKey      string `json:"adminApiKey"`
	AnalyticsApiKey  string `json:"analyticsApiKey"`
	MonitoringApiKey string `json:"monitoringApiKey"`
}

func NewCredentials(jsonData []byte) (Credentials, error) {
	var creds Credentials
	return creds, json.Unmarshal(jsonData, &creds)
}

func NewClientFromJSONAdmin(jsonData []byte) (algoliasearch.Client, error) {
	creds, err := NewCredentials(jsonData)
	if err != nil {
		return algoliasearch.NewClient("", ""), err
	}
	return algoliasearch.NewClient(creds.ApplicationId, creds.AdminApiKey), nil
}

func NewClientFromJSONSearchOrAdmin(jsonData []byte) (algoliasearch.Client, error) {
	creds, err := NewCredentials(jsonData)
	if err != nil {
		return algoliasearch.NewClient("", ""), err
	}
	if len(creds.SearchOnlyApiKey) > 0 {
		return algoliasearch.NewClient(creds.ApplicationId, creds.SearchOnlyApiKey), nil
	} else if len(creds.AdminApiKey) > 0 {
		return algoliasearch.NewClient(creds.ApplicationId, creds.AdminApiKey), nil
	}
	return algoliasearch.NewClient("", ""), errors.New("No Algolia Search or Admin API Key")
}
