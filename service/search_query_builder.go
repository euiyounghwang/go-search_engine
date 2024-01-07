package service

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/euiyounghwang/go-search_engine/lib/util"
	"github.com/euiyounghwang/go-search_engine/repository"
)


func add_highlighting() string {
	highlight_clauses := `
		"highlight": {
			"order": "score",
			"pre_tags": [
				"<b>"
			],
			"post_tags": [
				"</b>"
			],
			"fields": {
				"*": {
					"number_of_fragments": 1,
					"type": "plain",
					"fragment_size": 150
				}
			}
		}`
	
	return highlight_clauses
}



func Build_es_query(oas_query repository.Search) (string, error) {
	log.Println("Build_es_query..")
	
	// If no name was given, return an error with a message.
    // if oas_query == (repository.Search {}) {
    //     return "", errors.New("oas_query is empty")
    // }

	if reflect.ValueOf(oas_query).IsZero() {
        return "", errors.New("oas_query is empty")
    }
	
	s_source_clause := `
		"_source" : ["*"],
	`
	
	s_must_clause := ""
	if oas_query.Query_string == "" {
		s_must_clause = `{"match_all": {}}`
	} else {
		s_must_clause = `[
			{
				"query_string": {
					"fields": ["*"],
					"default_operator": "AND",
					"analyzer": "standard",
					"query": "%s"
				}
			}
		]`
	}
	
	// s_must_clause = fmt.Sprintf(s_must_clause, oas_query.Query_string)
	
	s_query_format := `
		"query" : {
			"bool" : {
				"must": %s,
				"should" : [],
				"filter": []
			}
		},
	`
	s_size_format := `
		"size" : %d,
	`
	es_query := fmt.Sprintf("{ %s %s %s %s }",
					s_source_clause,
					fmt.Sprintf(s_query_format, fmt.Sprintf(s_must_clause, oas_query.Query_string)),
					fmt.Sprintf(s_size_format, oas_query.Size),
					add_highlighting(),
				)
	
	log.Println("--")
	log.Printf("es_query : %s", util.PrettyString(es_query))
	log.Println("--")
	
	return es_query, nil
}