package main

import (
	// "go_swagger/docs"
	"go-search_engine/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	r := gin.Default()

	//코드로 SwaggerInfo 속성을 지정해지만 doc.json 정상적으로 조회된다.
	docs.SwaggerInfo.Title = "Golang Rest API"

    // 127.0.0.1:8080/docs/index.html 주로 swagger로 생성된 문서를 확인 수 있다. 
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// v1Group := r.Group("/api/v1")
	v1Group := r.Group("/")
	{
		v1Group.GET("/hello/:name", HelloHandler)
	}
	
	httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "9081"
    }
	r.Run("localhost:" + httpPort)
}

type Users struct {
	Id   int    `json:"id" example:"1"`      // UserId
	Name string `json:"name" example:"John"` // Name
	Age  int    `json:"age" example:"10"`    // Age
}

/* 아래 항목이 swagger에 의해 문서화 된다. */
// HelloHandler godoc
// @Summary test swagger api
// @Description test swagger api
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "Users name"
// @Router /hello/{name} [get]
// @Success 200 {object} Users
// @Failure 400
func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Users": ""})
	} else {
		user := Users{Id: 1, Name: name, Age: 20}
		c.JSON(http.StatusOK, gin.H{"Users": user})
	}
}