package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := initial_str
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return transform_str
}

func ReplaceStr(str string) (string) {
	var transformed_strg string = str
	transformed_strg = strings.Replace(transformed_strg, "\t\t", " ", -1)
	transformed_strg = strings.Replace(transformed_strg, " ", "", -1)
	transformed_strg = strings.Replace(transformed_strg, "\t", "", -1)
	return transformed_strg
}

func Uint8_to_Map(str []uint8) map[string]interface{}{
	var jsonMap map[string]interface{}
	// json.Unmarshal([]byte(str), &jsonMap)
	if err := json.Unmarshal([]byte(str), &jsonMap); err != nil {
        fmt.Println(err)
    }
	return jsonMap
}


func PrettyString(str string) (string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return ""
	}
	return prettyJSON.String()
}


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
	var sb strings.Builder
	_terms_array := strings.Split(_term, ",")
	for index, element := range _terms_array {
		sb.WriteString(`"` + element + `"`)
		if index != len(_terms_array)-1 {
			fmt.Println(index)
			sb.WriteString(`,`)		
		}
	}
	_terms_filters := `
	{
		"terms":{
			"_id":[
				%s
			]
		}
	}
	`
	_terms_filters_clause := fmt.Sprintf(_terms_filters, sb.String())
	fmt.Println("Build_terms_filters_batch - ", _terms_filters_clause)
	return _terms_filters_clause
}
