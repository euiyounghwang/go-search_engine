package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	// elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

var (
    IndexNameEmptyStringError = errors.New("index name cannot be empty string")
    IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
)


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := ""
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return replace_str
}

type ElasticDocs struct {
	SomeStr string
	SomeInt int
	SomeBool bool
	Timestamp int64
}

func main() {
	
	host := Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
    index := Set_Env(os.Getenv("ES_INDEX"), "go_test_omnisearch_v1")
	
	if index == "" {
        log.Fatal(IndexNameEmptyStringError)
    }

	es_client, err := elasticsearch.NewClient(
		elasticsearch.Config{
        	Addresses: []string{host},
			Username: "elastic",
			Password: "gsaadmin",
    	},
		// elastic.SetURL(host),
	)
	
	res, err := es_client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
		
	log.Println("Elasticsearch Initializing..")
	
	try_delete_index := func() {
		log.Println("inner func")
		response, err := es_client.Indices.Exists([]string{index})
		log.Println(response)
		if response.StatusCode != 404 {
			// log.Fatal(IndexAlreadyExistsError)
			es_client.Indices.Delete([]string{index})
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	
	try_delete_index()
	
    response, err := es_client.Indices.Create(index)
    if err != nil {
        log.Fatal(err)
    }

    // if response.IsError() {
    //     log.Fatal(err)
    // }
	
	if response.StatusCode == 200 {
		log.Println("Elasticsearch Indices Created..")
	}
	
	// Declare an empty slice for the Elasticsearch document struct objects
	docs := []ElasticDocs{}
	// Get the type of the 'docs' struct slice
	log.Println("docs TYPE:", reflect.TypeOf(docs))
	
	// New ElasticDocs struct instances
	newDoc1 := ElasticDocs{SomeStr: "Hello, world!", SomeInt: 42, SomeBool: true, Timestamp: 0.0}
	newDoc2 := ElasticDocs{SomeStr: "Hello, world2!", SomeInt: 7654, SomeBool: false, Timestamp: 0.0}
	
	// Append the new Elasticsearch document struct objects to the slice
	docs = append(docs, newDoc1)
	docs = append(docs, newDoc2)
	
	log.Println(docs)
	
	// res, _ = es_client.Index(index, esutil.NewJSONReader(&newDoc1))
	// fmt.Println(res)
	
	res, _ = es_client.Index(index, esutil.NewJSONReader(newDoc1))
	fmt.Println(res)
	defer res.Body.Close()
	
	// Deserialize the response into a map.
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		log.Printf("\nIndexRequest() RESPONSE:")
		// Print the response status and indexed document version.
		fmt.Println("Status:", res.Status())
		fmt.Println("Result:", resMap["_index"])
		// fmt.Println("Version:", int(resMap["_version"].(float64)))
		fmt.Println("resMap:", resMap)
		fmt.Println("\n")
	}
}