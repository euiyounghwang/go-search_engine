package util

import (
	"fmt"
	"strings"
)

/*
input :  "111", '222'
output :
[
	{
		"terms":{
			"_id":[
				"111"
			]
		}
	},
	{
		"terms":{
			"_id":[
				"222"
			]
		}
	}
]
*/
func Build_terms_filters_batch(_term string, _max_len int) string {
	
	_terms_filters_format := `
	{
		"bool": {
		  "must": [
			{
			  "bool": {
				"should": [
					%s
				]
			  }
			}
			]
		}
	}		  
	`
	
	var sb strings.Builder
	_terms_array := strings.Split(_term, ",")
	
	_terms_filters := `
	{
		"terms":{
			"_id":[
				%s
			]
		}
	}
	`
	for index, element := range _terms_array {
		if len(_terms_array) <= _max_len {
			sb.WriteString(`"` + element + `"`)
		} else {
			sb.WriteString(fmt.Sprintf(_terms_filters, `"` + element + `"`))
		}
		
		if index != len(_terms_array)-1 {
			sb.WriteString(`,`)		
		}	
	}
	
	var _terms_filters_clause string
	
	if len(_terms_array) <= _max_len {
		_terms_filters_clause = fmt.Sprintf(_terms_filters, sb.String())
	} else {
		_terms_filters_clause = sb.String()
	}
	
	well_formed_terms_filtered_batch := fmt.Sprintf(_terms_filters_format, _terms_filters_clause)
	fmt.Println("Build_terms_filters_batch - ",  PrettyString(string(well_formed_terms_filtered_batch)))
	return well_formed_terms_filtered_batch
}
