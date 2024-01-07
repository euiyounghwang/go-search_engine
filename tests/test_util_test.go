package test

import (
	"fmt"
	"os"
	"testing"

	util "github.com/euiyounghwang/go-search_engine/lib/util"

	"github.com/stretchr/testify/assert"
)

// go test -v ./tests/test_util_test.go

/*
-- Testing func

package hello

func Hello() string {
    return "Hello, world"
}

func TestHello(t *testing.T) {
	want := "Hello, world"
	fmt.Println(want)
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
*/

func Test_Get_Env(t *testing.T) {
	host := util.Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
	assert.Equal(t, host, "http://localhost:9209")

	os.Setenv("ES_HOST", "http://localhost:9200")
	assert.Equal(t, os.Getenv("ES_HOST"), "http://localhost:9200")
}



func Test_PrettyJSon(t *testing.T) {
	query := `{"track_total_hits" : true,"query": {"match_all" : {}},"size": 2}`
	
	var transformed_query_string string = util.PrettyString(query)
	fmt.Println(query)
	var expected_query string = `{
		"track_total_hits": true,
		"query": {
			"match_all": {}
		},
		"size": 2
	}`
	// assert.Equal(t, transformed_query_string, strings.Replace(expected_query, "\t\t", " ", -1))
	expected_query = util.ReplaceStr(expected_query)
	transformed_query_string = util.ReplaceStr(transformed_query_string)
	
	// fmt.Println(expected_query)
	assert.Equal(t, transformed_query_string, expected_query)
}

