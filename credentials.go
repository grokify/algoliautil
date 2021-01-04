package algoliautil

import (
	"encoding/json"
	"errors"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

// go get github.com/algolia/algoliasearch-client-go/v3@v3.14.0

const ENV_VAR_APP_CREDENTIALS_JSON = "ALGOLIA_APP_CREDENTIALS_JSON"

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

func NewClientFromJSONAdmin(jsonData []byte) (*search.Client, error) {
	creds, err := NewCredentials(jsonData)
	if err != nil {
		return nil, err
	}
	return search.NewClient(creds.ApplicationId, creds.AdminApiKey), nil
}

func NewClientFromJSONSearchOrAdmin(jsonData []byte) (*search.Client, error) {
	creds, err := NewCredentials(jsonData)
	if err != nil {
		return nil, err
	}
	if len(creds.AdminApiKey) > 0 {
		return search.NewClient(creds.ApplicationId, creds.AdminApiKey), nil
	} else if len(creds.SearchOnlyApiKey) > 0 {
		return search.NewClient(creds.ApplicationId, creds.SearchOnlyApiKey), nil
	}
	return search.NewClient("", ""), errors.New("No Algolia Search or Admin API Key")
}
