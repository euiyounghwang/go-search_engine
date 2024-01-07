package repository

import "errors"

// https://transform.tools/json-to-go

var (
    IndexNameEmptyStringError = errors.New("index name cannot be empty string")
    IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
	ESInstanceError   		  = errors.New("elasticsearch goes down")
)

type Search struct {
	IdsFilter        []string `json:"ids_filter" example:"performance"`
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


type Search_Results struct {
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Hits struct {
		Hits []struct {
			ID     string  `json:"_id"`
			Index  string  `json:"_index"`
			Score  float64 `json:"_score"`
			Source struct {
				Timestamp       string  `json:"@timestamp"`
				ConcurrentUsers string  `json:"concurrent_users"`
				ElapsedTime     float64 `json:"elapsed_time"`
				EntityType      string  `json:"entity_type"`
				Env             string  `json:"env"`
				SearchIndex     string  `json:"search_index"`
				Sequence        int     `json:"sequence"`
				Title           string  `json:"title"`
			} `json:"_source"`
			Highlight struct {
				EntityType   []string `json:"entity_type"`
				Title        []string `json:"title"`
				TitleKeyword []string `json:"title.keyword"`
			} `json:"highlight"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	TimedOut bool `json:"timed_out"`
	Took     int  `json:"took"`
}