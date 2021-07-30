package algoliautil

import (
	"encoding/json"
	"errors"
	"strings"

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

func (c Credentials) TrimSpace() {
	c.ApplicationId = strings.TrimSpace(c.ApplicationId)
	c.SearchOnlyApiKey = strings.TrimSpace(c.SearchOnlyApiKey)
	c.AdminApiKey = strings.TrimSpace(c.AdminApiKey)
	c.AnalyticsApiKey = strings.TrimSpace(c.AnalyticsApiKey)
	c.MonitoringApiKey = strings.TrimSpace(c.MonitoringApiKey)
}

func NewClient(c Credentials) (*search.Client, error) {
	c.TrimSpace()
	if len(c.ApplicationId) == 0 {
		return nil, errors.New("no Algolia applicationId")
	}
	if len(c.AdminApiKey) > 0 {
		return search.NewClient(c.ApplicationId, c.AdminApiKey), nil
	} else if len(c.SearchOnlyApiKey) > 0 {
		return search.NewClient(c.ApplicationId, c.SearchOnlyApiKey), nil
	}
	return nil, errors.New("no Algolia Search or Admin API Key")
}

func NewCredentials(jsonData []byte) (Credentials, error) {
	var creds Credentials
	return creds, json.Unmarshal(jsonData, &creds)
}

func NewClientJSON(jsonData []byte) (*search.Client, error) {
	creds, err := NewCredentials(jsonData)
	if err != nil {
		return nil, err
	}
	return NewClient(creds)
}
