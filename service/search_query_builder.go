package service

import "fmt"


func Build_es_query(oas_query string) string {
	s_must_clause := `[
		{
			"query_string": {
				"fields": ["*"],
				"default_operator": "AND",
				"analyzer": "standard",
				"query": "%s"
			}
		}
	]`
	s_must_clause = fmt.Sprintf(s_must_clause, "performance")
	
	s_query_format := `{
		"query" : {
			"bool" : {
				"must": %s,
				"should" : [],
				"filter": []
			}
		},
		"size": 2
	}`
	
	es_query := fmt.Sprintf(s_query_format, s_must_clause,)
	
	return es_query
}