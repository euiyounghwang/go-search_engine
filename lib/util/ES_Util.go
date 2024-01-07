package util

import (
	"fmt"
	"strings"
)

/*
INPUT : "111, 222"
OUTPUT : "111", "222"
*/
func Build_split_string_array(s string) string {
	var sb strings.Builder
	s_array := strings.Split(s, ",")
	
	for index, element := range s_array {
		sb.WriteString(`"` + element + `"`)
		if index != len(s_array)-1 {
			sb.WriteString(`,`)		
		}	
	}
	
	return sb.String()
}

/*
input :  "111, 222"
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
func Build_terms_filters_batch(_term string, max_terms_count int) string {
	
	if _term == "" {
		_terms_filters_nothing :=  `[]`
		// fmt.Printf("Build_terms_filters_batch [_term : %s, max_terms_count : %d] - %s\n",  
		// 	_term,
		// 	max_terms_count, 
		// 	PrettyString(_terms_filters_nothing),
		// )
		return _terms_filters_nothing
	}
	
	if max_terms_count < 1 {
		max_terms_count = 65000
	}
	
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
		if len(_terms_array) <= max_terms_count {
			sb.WriteString(`"` + element + `"`)
		} else {
			sb.WriteString(fmt.Sprintf(_terms_filters, `"` + element + `"`))
		}
		
		if index != len(_terms_array)-1 {
			sb.WriteString(`,`)		
		}	
	}
	
	var _terms_filters_clause string
	if len(_terms_array) <= max_terms_count {
		_terms_filters_clause = fmt.Sprintf(_terms_filters, sb.String())
	} else {
		_terms_filters_clause = sb.String()
	}
	
	_terms_filters_format := `[
		%s
	]`
	
	well_formed_terms_filtered_batch := fmt.Sprintf(_terms_filters_format, _terms_filters_clause)
	fmt.Printf(
		"Build_terms_filters_batch [_term : %s, max_terms_count : %d] - %s\n",  
		_term, 
		max_terms_count, 
		PrettyString(string(well_formed_terms_filtered_batch)),
	)
	
	return well_formed_terms_filtered_batch
}
