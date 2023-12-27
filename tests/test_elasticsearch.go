package main

import (
	"testing"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func Test_elastic(t *testing.T) {
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
        	Addresses: []string{es_host},
			Username: "elastic",
			Password: "gsaadmin",
    	},
		// elastic.SetURL(host),
	)
	if err != nil {
		panic(err)
	}

    res, err := es.Info()
    if err != nil {
     panic(err)
    }
    defer res.Body.Close()
}