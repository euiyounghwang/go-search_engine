package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
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
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body) ), &jsonMap)
	
	log.Println(jsonMap)

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Error closing response body: %v", err)
		}
	}()
}