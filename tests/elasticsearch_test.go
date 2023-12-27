package test

import (
	"testing"

	// elasticsearch "github.com/elastic/go-elasticsearch/v8"
	my_elasticsearch "go-search_engine/tools/lib"
)

/*
go test  -v
go test *_test.go -v
Specific TestFunctionName) $ go test -run TestFunctionName -v
*/

func Test_elastic(t *testing.T) {
	es_host := "http://localhost:9209"
	es := my_elasticsearch.Get_es_instance(es_host)
    res, err := es.Info()
    if err != nil {
     panic(err)
    }
    defer res.Body.Close()
}