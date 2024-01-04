package controller

import (
	"go-search_engine/repository"
	"go-search_engine/service"
	"log"
	"net/http"
	"reflect"
	"time"

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
	
	es_host, _ := c.Get("ES_HOST")
	log.Println(es_host, reflect.TypeOf(es_host))
	
	response_json := service.Build_ES_Instance_Health(es_host.(string))
	
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
	
	// r := &search
	// log.Println(search.Query_string, search.Sort_order, r.Query_string)
	
	// if es_host, exists := c.Get("ES_HOST"); exists {
	// 	print("exists", es_host, exists)
	// }
	es_host, _ := c.Get("ES_HOST")
	index_name, _ := c.Get("Index_Name")
	log.Println(es_host, reflect.TypeOf(es_host), search, reflect.TypeOf(search))
		
	response_json := service.Build_search(es_host.(string), index_name.(string), search)
	
	// End Time request
    endTime := time.Now()
	// execution time
    latencyTime := endTime.Sub(startTime)

	log.Printf("Excuting Time :  %s", latencyTime)

	// c.JSON(http.StatusOK, gin.H{"message": "success"})
	c.JSON(http.StatusOK, response_json)
}


