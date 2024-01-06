package main

import (
	// "go_swagger/docs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/euiyounghwang/go-search_engine/docs"

	"github.com/euiyounghwang/go-search_engine/lib/util"

	controller "github.com/euiyounghwang/go-search_engine/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Transform json to struct in Go, https://transform.tools/json-to-go
type Config struct {
	App struct {
		Es struct {
			EsHost string `json:"es_host"`
			Index  struct {
				Alias string `json:"alias"`
			} `json:"index"`
		} `json:"es"`
	} `json:"app"`
}

func read_yml() Config {
	
	yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
    	fmt.Printf("yamlFile.Get err #%v ", err)
    }
	
	config := Config{}
    if err := json.Unmarshal(yamlFile, &config); err != nil {
        // do error check
        fmt.Println(err)
    }

	fmt.Println("--")
	fmt.Println("read_yml()")
    fmt.Println(config)
	// fmt.Println(config.App.Es.EsHost, config.App.Es.Index.Alias)
	fmt.Println("--")
	
	// fmt.Println(obj["app"].(map[interface {}]interface{})["es"].(map[interface {}]interface{})["es_host"])
	return config
}


func init_params(c *gin.Context) {
	config := read_yml()
	c.Set("ES_HOST", util.Set_Env(os.Getenv("ES_HOST"), config.App.Es.EsHost))
	c.Set("Index_Name",config.App.Es.Index.Alias)
}

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

	// read_yml()
	
	// https://github.com/swaggo/swag/blob/master/README.md
	docs.SwaggerInfo.Title = "Golang Rest API"
	
    // 127.0.0.1:8080/docs/index.html 주로 swagger로 생성된 문서를 확인 수 있다. 
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	  })
	

	// v1Group := r.Group("/api/v1")
	v1Group := r.Group("/")
	{
		v1Group.GET("/", defaultHandler)
		v1Group.GET("/hello/:name", HelloHandler)
	}
	
	v1Search := r.Group("/")
	{
		v1Search.GET("/health", init_params, controller.HealthHandler)
		v1Search.POST("/es/search", init_params, controller.SearchHandler)
	}
	
	httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "9081"
    }
	r.Run("0.0.0.0:" + httpPort)
}

type Users struct {
	Id   int    `json:"id" example:"1"`      // UserId
	Name string `json:"name" example:"John"` // Name
	Age  int    `json:"age" example:"10"`    // Age
}

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @tags API
// HelloHandler godoc
// @Summary test swagger api
// @Description test swagger api
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router / [get]
// @Success 200 
// @Failure 400
func defaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
}

/*
// SearchHandler godoc
// @tags API
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

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @tags API
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
