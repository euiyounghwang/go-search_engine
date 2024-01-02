package repository

import "errors"


var (
    IndexNameEmptyStringError = errors.New("index name cannot be empty string")
    IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
	ESInstanceError   		  = errors.New("elasticsearch goes down")
)

type Search struct {
	Include_basic_aggs bool `json:"include_basic_aggs" example:"true"` 
	Pit string `json:"pit" example:""`
	Query_string string `json:"query_string" example:"performance"`
	Size   int    `json:"size" example:"10"`
	Sort_order string `json:"sort_order" example:"DESC"`
	Start_date string `json:"start_date" example:"2021 01-01 00:00:00"`
}	

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}