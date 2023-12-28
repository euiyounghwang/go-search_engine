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

var es_host = my_elasticsearch.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
var es_client = my_elasticsearch.Get_es_instance(es_host)
// es := my_elasticsearch.Get_es_instance(es_host)

func Test_elasticsearch_instance_status(t *testing.T) {
	
	// es_client := my_elasticsearch.Get_es_instance(es_host)
	
	// es is not None
	assert.Equal(t, es_client == nil, false)
	
    res, err := es_client.Info()
	log.Println(res)
    if err != nil {
     panic(err)
    }
	defer res.Body.Close()
	
	assert.Equal(t, res.StatusCode, 200)
}



func Test_elasticsearch_setup(t *testing.T) {
	
	index := "test_omnisearch_v1_go"
	// es_client := my_elasticsearch.Get_es_instance(es_host)
	
	assert.Equal(t, es_client != nil, true)
	
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
	
	try_create_index := func() {
		res, err := es_client.Indices.Create(index)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
		assert.Equal(t, res.StatusCode, 200)
	}
	
	try_create_alias := func() {
		
	}
	
	try_delete_index()
	try_create_index()
	try_create_alias()
	
}