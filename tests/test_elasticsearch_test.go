package test

import (
	"bytes"
	"context"
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
		
	create_alias := func(index string, alias string) {
		
		type UpdateAliasAction struct {
			Index string `json:"index"`
			Alias string `json:"alias"`
		}
		type UpdateAliasRequest struct {
			Actions []map[string]*UpdateAliasAction `json:"actions"`
		}
		
		updateActions := make([]map[string]*UpdateAliasAction, 0)
		addAction := make(map[string]*UpdateAliasAction)
		addAction["add"] = &UpdateAliasAction{
			Index: index,
			Alias: alias,
		}
		updateActions = append(updateActions, addAction)
		jsonBody, err := json.Marshal(&UpdateAliasRequest{
			Actions: updateActions,
		})
		if err != nil {
			log.Fatal(err)
		}
		
		// make API request
		res, err := es_client.Indices.UpdateAliases(
			bytes.NewBuffer(jsonBody),
			es_client.Indices.UpdateAliases.WithContext(context.Background()),
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
	}
	create_alias("test_performance_metrics_v1", "metrics_search")
	
	Index_with_document := func(index string) {
		/*
		-- Python
		es.index(index="test_performance_metrics_v1", id=111, body={
			"title" :  "performance",
			"elapsed_time": 0.3,
			"sequence": 1,
			"entity_type": "performance",
			"env" :  "dev",
			"concurrent_users" :  "20",
			"search_index" :  "test_performance_metrics_v1",
			"@timestamp" : "2023-01-01 00:00:00"
			}
		)
		*/
		res, err := es_client.Index(
			index,                               // Index name
			strings.NewReader(`{
				"title" :  "performance",
				"elapsed_time": 0.3,
				"sequence": 1,
				"entity_type": "performance",
				"env" :  "dev",
				"concurrent_users" :  "20",
				"search_index" :  "test_performance_metrics_v1",
				"@timestamp" : "2023-01-01 00:00:00"
				}`), // Document body
			es_client.Index.WithDocumentID("1"),            // Document ID
			es_client.Index.WithRefresh("true"),            // Refresh
		)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		body, _ := io.ReadAll(res.Body)
		log.Printf("Index_with_document - [%s]", util.PrettyString(string(body)))
		defer res.Body.Close()
	}
	
	Index_with_document("test_performance_metrics_v1")
}