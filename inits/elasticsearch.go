package inits

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/yuanzhichao-star/public_pkg/config"

	"log"
)

func InitEs() {
	var err error
	es := config.AppCong.Elasticsearch
	addr := fmt.Sprintf("%s:%d", es.Host, es.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			addr,
		},
	}
	config.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("es init success")
}
