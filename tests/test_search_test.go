package test

import (
	"testing"

	"github.com/euiyounghwang/go-search_engine/lib/util"
	"github.com/stretchr/testify/assert"
)

// go test -v ./tests/test_search_test.go

func Test_terms_filters_batch(t *testing.T) {
	/*
	returns_terms_filters :=
	{
	"terms":{
	"_id":[
	"111","222"
	]
	}
	}
	*/
	ids_filter := "111,222"
	
	_max_len := 2
	returns_terms_filters := util.Build_terms_filters_batch(ids_filter, _max_len)
	expected_terms_filters := `
	[
		{
			"terms": {
				"_id": [
					"111",
					"222"
				]
			}
		}
	]
	`
	assert.Equal(t, util.PrettyString(util.ReplaceStr(returns_terms_filters)), util.PrettyString(util.ReplaceStr(expected_terms_filters)))
	
	_max_len = 0
	returns_terms_filters = util.Build_terms_filters_batch(ids_filter, _max_len)
	assert.Equal(t, util.PrettyString(util.ReplaceStr(returns_terms_filters)), util.PrettyString(util.ReplaceStr(expected_terms_filters)))
	
	_max_len = 1
	returns_terms_filters = util.Build_terms_filters_batch(ids_filter, _max_len)
	expected_terms_filters = `
	[
		{
			"terms": {
				"_id": [
					"111"
				]
			}
		},
		{
			"terms": {
				"_id": [
					"222"
				]
			}
		}
	]   
	`
	assert.Equal(t, util.PrettyString(util.ReplaceStr(returns_terms_filters)), util.PrettyString(util.ReplaceStr(expected_terms_filters)))
	
	
	
	ids_filter = ""
	_max_len = 2
	returns_terms_filters = util.Build_terms_filters_batch(ids_filter, _max_len)
	expected_terms_filters = `[]`
	assert.Equal(t, util.PrettyString(util.ReplaceStr(returns_terms_filters)), util.PrettyString(util.ReplaceStr(expected_terms_filters)))
}