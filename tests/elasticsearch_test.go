package test

import (
	"log"
	"os"
	"testing"

	// elasticsearch "github.com/elastic/go-elasticsearch/v8"
	my_elasticsearch "go-search_engine/lib"

	"github.com/stretchr/testify/assert"
	// my_elasticsearch "github.com/euiyounghwang/go-search_engine/lib"
)

/*
go test  -v
go test *_test.go -v
Specific TestFunctionName) $ go test -run TestFunctionName -v
*/

var es_host string = my_elasticsearch.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
// var es *elasticsearch.Client

func Test_elasticsearch_instance_status(t *testing.T) {
	
	es := my_elasticsearch.Get_es_instance(es_host)
    res, err := es.Info()
	log.Println(res)
    if err != nil {
     panic(err)
    }
    defer res.Body.Close()
}



func Test_elasticsearch_setup(t *testing.T) {
	index := "test_omnisearch_v1_go"
	es_client := my_elasticsearch.Get_es_instance(es_host)
	try_delete_index := func() {
		response, err := es_client.Indices.Exists([]string{index})
		log.Println(response)
		if response.StatusCode != 404 {
			es_client.Indices.Delete([]string{index})
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	
	try_delete_index()
	
	res, err := es_client.Indices.Create(index)
    if err != nil {
        log.Fatal(err)
    }
	log.Println(res)
	assert.Equal(t, res.StatusCode, 200)
}