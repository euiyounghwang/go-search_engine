package repository

import "errors"


var (
    IndexNameEmptyStringError = errors.New("index name cannot be empty string")
    IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
	ESInstanceError   		  = errors.New("elasticsearch goes down")
)

type Search struct {
	Id   int    `json:"id" example:"1"`      // UserId
	Name string `json:"name" example:"John"` // Name
	Age  int    `json:"age" example:"10"`    // Age
}


type APIError struct {
	ErrorCode    int
	ErrorMessage string
}