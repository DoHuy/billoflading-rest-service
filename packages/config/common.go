package config

import (
	"github.com/elastic/go-elasticsearch/v6"
	"net/http"
	"strings"
	"time"
)

type HTTPConfig struct {
	Addr string `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:8080"`
}
type RunModeConfig struct {
	RunMode string `envconfig:"RUN_MODE" required:"true"`
}
type ElasticConfig struct {
	ElasticURI string `envconfig:"ELASTIC_URI" required:"true"`
}
type OracleConfig struct {
	OracleURI string `envconfig:"ORACLE_URI" required:"true"`
}

func InitElasticClient(es ElasticConfig) (*elasticsearch.Client, error) {
	addrs := strings.Split(es.ElasticURI, ",")
	cfg := elasticsearch.Config{
		Addresses: addrs,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
		},
	}
	return elasticsearch.NewClient(cfg)
}

//func InitOracleClient(or OracleConfig) () {
//
//}
