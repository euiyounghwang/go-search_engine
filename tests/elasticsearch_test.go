package test

import (
	"testing"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

/*
go test  -v
go test *_test.go -v
Specific TestFunctionName) $ go test -run TestFunctionName -v
*/

func Test_elastic(t *testing.T) {
	es_host := "http://localhost:9209"
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