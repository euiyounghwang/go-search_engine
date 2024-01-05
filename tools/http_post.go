package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"runtime"

	"github.com/euiyounghwang/go-search_engine/repository"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}


func Struct_Iterate_Rows(body []uint8) {
	/*
	Iterate body using Struct for search results
	*/
	
	fmt.Println("--")
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body) ), &jsonMap)
	log.Println(jsonMap)
	fmt.Println("--")
	fmt.Println()
	
	// using Struct
	result := repository.Search_Results{}
	if err := json.Unmarshal(body, &result); err != nil {
        // do error check
        fmt.Println(err)
    }
	
	fmt.Println("--")
	// fmt.Println((result))
	log.Printf("The number of documents: %d", result.Hits.Total.Value)
	
	fmt.Printf("[%s func] result.Hits.Hits", runtime.FuncForPC(reflect.ValueOf(Struct_Iterate_Rows).Pointer()).Name())
	for i, rows := range result.Hits.Hits {
		fmt.Println("sequence : ", i+1)
		fmt.Println("rows.Source.Title : ", rows.Source.SearchIndex)
		fmt.Println("rows.Source.Title : ", rows.Source.SearchIndex)
	}
	fmt.Println("--")
	fmt.Println()
}


func main() {
	posturl := "http://localhost:9081/es/search"

	data := []byte(`
	{
		"include_basic_aggs": true,
		"pit": "",
		"query_string": "performance",
		"size": 10,
		"sort_order": "DESC",
		"start_date": "2021 01-01 00:00:00"
	  }
	`)

	// {"index": {"_index": "test_performance_metrics_v1", "_id": "2"}}
	// {"title" :  "performance", "elapsed_time": 0.3, "sequence": 1, "entity_type": "performance", "env" :  "dev", "concurrent_users" :  "20", "search_index" :  "test_performance_metrics_v1", "@timestamp" : "2023-01-01 00:00:00"}
	
	// Create a new request using http
	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// add authorization header to the request
	// req.Header.Add("Authorization", "Bearer $API_TOKEN")
	req.Header.Add("Content-Type", "application/json")

	// send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error on response: %v", err)
	}
	
	// log.Println(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	
	Struct_Iterate_Rows(body)
	
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Error closing response body: %v", err)
		}
	}()
}