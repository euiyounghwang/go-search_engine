package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := ""
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return replace_str
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
