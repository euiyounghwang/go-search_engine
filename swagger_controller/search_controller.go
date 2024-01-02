package controller

import (
	"go-search_engine/repository"
	"log"
	"net/http"
	"os"
	"reflect"
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
// @Summary search engine api
// @tags Search
// @Description search engine api
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
	
	log.Println(res)
	c.JSON(http.StatusOK, gin.H{"message": res})
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
	c.JSON(http.StatusOK, gin.H{"message": search})
}


