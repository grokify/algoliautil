package algoliautil

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

// go get github.com/algolia/algoliasearch-client-go/v3@v3.14.0

const EnvVarAppCredentialsJSON = "ALGOLIA_APP_CREDENTIALS_JSON"

type Credentials struct {
	ApplicationID    string `json:"applicationId"`
	SearchOnlyAPIKey string `json:"searchOnlyApiKey"`
	AdminAPIKey      string `json:"adminApiKey"`
	AnalyticsAPIKey  string `json:"analyticsApiKey"`
	MonitoringAPIKey string `json:"monitoringApiKey"`
}

func (c Credentials) TrimSpace() {
	c.ApplicationID = strings.TrimSpace(c.ApplicationID)
	c.SearchOnlyAPIKey = strings.TrimSpace(c.SearchOnlyAPIKey)
	c.AdminAPIKey = strings.TrimSpace(c.AdminAPIKey)
	c.AnalyticsAPIKey = strings.TrimSpace(c.AnalyticsAPIKey)
	c.MonitoringAPIKey = strings.TrimSpace(c.MonitoringAPIKey)
}

func NewClient(c Credentials) (*search.Client, error) {
	c.TrimSpace()
	if len(c.ApplicationID) == 0 {
		return nil, errors.New("no Algolia applicationId")
	}
	if len(c.AdminAPIKey) > 0 {
		return search.NewClient(c.ApplicationID, c.AdminAPIKey), nil
	} else if len(c.SearchOnlyAPIKey) > 0 {
		return search.NewClient(c.ApplicationID, c.SearchOnlyAPIKey), nil
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
