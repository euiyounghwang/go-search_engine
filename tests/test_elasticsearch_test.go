package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	// elasticsearch "github.com/elastic/go-elasticsearch/v8"
	my_elasticsearch "go-search_engine/lib/elasticsearch"
	util "go-search_engine/lib/util"

	"github.com/stretchr/testify/assert"
	// my_elasticsearch "github.com/euiyounghwang/go-search_engine/lib"
)

/*
go test  -v
go test *_test.go -v
Specific TestFunctionName) $ go test -run TestFunctionName -v
*/

var es_host = util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
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



func Test_elasticsearch_configuration_to_local(t *testing.T) {
	
	// index := "test_omnisearch_v1_go"
	// es_client := my_elasticsearch.Get_es_instance(es_host)
	
	assert.Equal(t, es_client != nil, true)
	
	try_delete_index := func(index string) {
		response, err := es_client.Indices.Exists([]string{index})
		log.Println(response)
		
		assert.Equal(t, response.StatusCode, 200)
		
		if response.StatusCode != 404 {
			res, _ := es_client.Indices.Delete([]string{index})
			assert.Equal(t, res.StatusCode, 200)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	try_delete_index("test_ngram_v1")
	try_delete_index("test_performance_metrics_v1")
	
	try_create_index := func(index string, filelname string) {
		mapping_json, err := os.Open(filelname)
		if err != nil {
			log.Fatal(err)
		}
		defer mapping_json.Close()
		
		// fmt.Println(mapping_json)
		
		mapping_json_raw, err := ioutil.ReadAll(mapping_json)
		if err != nil {
			fmt.Printf("failed to read json file, error: %v", err)
			return
		}
		// fmt.Println(mapping_json_raw)
		
		mapping := make(map[string]interface{})
		if err := json.Unmarshal(mapping_json_raw, &mapping); err != nil {
			log.Fatal(err)
		}
		
		/*
		mapping_json_raw := `
		{
		"settings": {
			"number_of_shards": 1
		},
		"mappings": {
			"properties": {
			"field1": {
				"type": "text"
			}
			}
		}
		}`
		*/
		// fmt.Println(mapping["settings"])
		
		res, err := es_client.Indices.Create(index, es_client.Indices.Create.WithBody(strings.NewReader(string(mapping_json_raw))))
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(res)
		
		body, _ := io.ReadAll(res.Body)
		log.Printf("try_create_index - [%s]", util.PrettyString(string(body)))
		res.Body.Close()
		
		log.Println("type ", reflect.TypeOf(body))
		response_json := util.StringJson_to_Json(body)
		log.Printf("Json : %s, parsing : [%s]", response_json, response_json["index"])
			
		assert.Equal(t, res.StatusCode, 200)
	}
	try_create_index("test_ngram_v1", "./test_mapping/performance_metrics_mapping.json")
	try_create_index("test_performance_metrics_v1", "./test_mapping/search_ngram_mapping.json")
		
	try_create_alias := func() {
		
	}
	try_create_alias()
	
}