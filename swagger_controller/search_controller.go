package controller

import (
	"context"
	"go-search_engine/repository"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	my_elasticsearch "go-search_engine/lib/elasticsearch"
	"go-search_engine/lib/util"

	"github.com/gin-gonic/gin"
)

/*
// SearchHandler godoc
// @Summary search engine api
// @tags Search
// @Description search engine api
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /es/search [get]
// @Success 200
// @Failure 400
func SearchHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
}
*/

// HealthHandler godoc
// @Summary search engine health
// @tags Search
// @Description search engine health
// @Accept  json
// @Produce  json
// @Router /health [get]
// @content application/json
// @Success 200 {object} object
// @Failure 500 {object} object
func HealthHandler(c *gin.Context) {
	
	es_host := util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
	es_client := my_elasticsearch.Get_es_instance(es_host)
	
	res, err := es_client.Info()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": repository.ESInstanceError.Error()})
		log.Printf("Error getting response: %s", err)
		return
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
	
	// c.JSON(http.StatusOK, gin.H{"message": response_json})
	c.JSON(http.StatusOK, response_json)
}

// SearchHandler godoc
// @Summary search engine api
// @tags Search
// @Description search engine api
// @Accept  json
// @Produce  json
// @Router /es/search [post]
// @Param search body repository.Search true "Search Info Body"
// @content application/json
// @Success 200 {object} repository.Search
// @Failure 400,404,500 {object} object
/*
// @Failure 400 {object} repository.APIError "We need ID!!"
// @Failure 404 {object} repository.APIError "Can not find ID"
*/
func SearchHandler(c *gin.Context) {
	// Starting time request
    startTime := time.Now()
	
	var search repository.Search
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// End Time request
    endTime := time.Now()
	// execution time
    latencyTime := endTime.Sub(startTime)

	// c.JSON(http.StatusOK, gin.H{"message": "success"})
	log.Printf("Parsing :  %s", reflect.TypeOf(search))
	log.Printf("Excuting Time :  %s", latencyTime)
	
	
	es_host := util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
	es_client := my_elasticsearch.Get_es_instance(es_host)
	
	query := `{
		"track_total_hits" : true,
		"query": {
			"match_all" : {}
		},
		"size": 2
	}`
	
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
	response_json := util.Uint8_to_Map(body)
	
	c.JSON(http.StatusOK, response_json)
}


