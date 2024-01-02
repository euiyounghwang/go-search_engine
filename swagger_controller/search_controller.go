package swagger_search

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Search struct {
	Id   int    `json:"id" example:"1"`      // UserId
	Name string `json:"name" example:"John"` // Name
	Age  int    `json:"age" example:"10"`    // Age
}


type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

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

// SearchHandler godoc
// @Summary search engine api
// @tags Search
// @Description search engine api
// @Accept  json
// @Produce  json
// @Router /es/search [post]
// @Param search body Search true "Search Info Body"
// @content application/json
// @Success 200 {object} Search
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
func SearchHandler(c *gin.Context) {
	var search Search
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"message": "success"})
	c.JSON(http.StatusOK, gin.H{"message": search})
}
