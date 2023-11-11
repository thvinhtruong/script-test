package api

import (
	"fmt"
)

const (
	httpRequestPrefix = "http://localhost:9000/api/v1/GetUserRecord/"
)

type APIConfig struct {
	Endpoint    string `json:"endpoint"`
	EnableCache bool   `json:"enablecache"`
}

func NewAPIConfig(enableCache bool) *APIConfig {
	return &APIConfig{
		Endpoint:    httpRequestPrefix,
		EnableCache: enableCache,
	}
}

func (c *APIConfig) SetAPIEndpoint(userId string) {
	c.Endpoint = fmt.Sprintf(httpRequestPrefix + userId)
	if c.EnableCache {
		c.Endpoint += "?cacheEnable=1"
	} else {
		c.Endpoint += "?cacheEnable=0"
	}
}

func (c *APIConfig) GetAPIEndpoint() string {
	return c.Endpoint
}
