package service

import (
	"context"
	my_elasticsearch "go-search_engine/lib/elasticsearch"
	"go-search_engine/lib/util"
	"go-search_engine/repository"
	"io"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func Build_search(es_host string, query string) map[string]interface{} {
	// es_host := util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
	es_client := my_elasticsearch.Get_es_instance(es_host)
	
	ctx := context.Background()
	
	// var b strings.Builder
	// b.WriteString(query)
	// read := strings.NewReader(b.String())
    res, err := es_client.Search(
		es_client.Search.WithContext(ctx),
		es_client.Search.WithIndex("test_performance_metrics_v1"),
		es_client.Search.WithBody(strings.NewReader(query)),
		es_client.Search.WithTrackTotalHits(true),
		es_client.Search.WithPretty(),
		es_client.Search.WithFrom(0),
		es_client.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	log.Println(res)
	defer res.Body.Close()
	
	body, _ := io.ReadAll(res.Body)
	
	// var response_json map[string]interface{}
	// json.Unmarshal([]byte(string(body) ), &response_json)
	
	response_json := util.Uint8_to_Map(body)
	
	return response_json
	
}


func Build_ES_Instance_Health(es_host string) map[string]interface{} {
	// es_host := util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
	es_client := my_elasticsearch.Get_es_instance(es_host)
	
	res, err := es_client.Info()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": repository.ESInstanceError.Error()})
		log.Printf("Error getting response: %s", err)
		// return
		return gin.H{"message": repository.ESInstanceError.Error()}
	}
	
	if res.StatusCode == 200 {
		log.Println(res, reflect.TypeOf(res))
	}
	
	body, _ := io.ReadAll(res.Body)
	response_json := util.Uint8_to_Map(body)
	log.Println("Uint8_to_Map type - ", reflect.TypeOf(response_json))
	log.Printf("Json : %s, parsing : %s, %s", response_json, response_json["cluster_name"])
	
	// var jsonMap map[string]interface{}
	// json.Unmarshal([]byte(string(body) ), &jsonMap)
	return response_json
}