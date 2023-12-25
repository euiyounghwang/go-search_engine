package main

import (
	"errors"
	"log"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
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

func main() {
	
	host := Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
    index := Set_Env(os.Getenv("ES_INDEX"), "go_test_omnisearch_v1")
	
	if index == "" {
        log.Fatal(IndexNameEmptyStringError)
    }

    es_client, err := elasticsearch.NewClient(elasticsearch.Config{
        Addresses: []string{host},
    })
	
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
}