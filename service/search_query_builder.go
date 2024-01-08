package service

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/euiyounghwang/go-search_engine/lib/util"
	"github.com/euiyounghwang/go-search_engine/repository"
)



func build_source_clause(oas_query repository.Search) string {
	var s_source_clause string
	if oas_query.Source_fields == "" || oas_query.Source_fields == "*" {
		s_source_clause = `"*"`
	} else {
		s_source_clause = util.Build_split_string_array(oas_query.Source_fields)
	}
	
	return s_source_clause
}

func build_must_clauses(oas_query repository.Search) string {
	s_must_clause := ""
	if oas_query.Query_string == "" {
		s_must_clause = `{"match_all": {}}`
	} else {
		s_must_clause = `
		{
				"query_string": {
					"fields": ["*"],
					"default_operator": "AND",
					"analyzer": "standard",
					"query": "%s"
				}
		}`
		s_must_clause = fmt.Sprintf(s_must_clause, oas_query.Query_string)
		
	}
	
	return s_must_clause
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
	
	// Build Source Clause
	s_source_clause := build_source_clause(oas_query)
	// Build must Clause
	s_must_clause := build_must_clauses(oas_query)
	
	es_query_format := `
	{
		"_source" : [%s],
		"track_total_hits": true,
		"query" : {
			"bool" : {
				"must": [%s],
				"should" : [],
				"filter": [
					{
						"bool": {
							"must": [
								{
									"bool": {
										"should": %s
									}
								}
							]
						}
					}
				]
			}
		},
		"size" : %d,
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
		}
	}
	`
	es_query := fmt.Sprintf(es_query_format,
						   s_source_clause,
						   s_must_clause, 
						   util.Build_terms_filters_batch(oas_query.IdsFilter, -1),
						   oas_query.Size,
						  )
	
	log.Println("--")
	log.Printf("es_query : %s", util.PrettyString(es_query))
	log.Println("--")
	
	return es_query, nil
}