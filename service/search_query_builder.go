package service

import (
	"errors"
	"fmt"
	"go-search_engine/lib/util"
	"go-search_engine/repository"
	"log"
)


func Build_es_query(oas_query repository.Search) (string, error) {
	log.Println("Build_es_query..")
	
	// If no name was given, return an error with a message.
    if oas_query == (repository.Search {}) {
        return "", errors.New("oas_query is empty")
    }

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
	s_must_clause = fmt.Sprintf(s_must_clause, oas_query.Query_string)
	
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
	
	log.Println("--")
	log.Printf("es_query : %s", util.PrettyString(es_query))
	log.Println("--")
	
	return es_query, nil
}