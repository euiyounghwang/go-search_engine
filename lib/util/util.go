package util

import (
	"bytes"
	"encoding/json"
	"log"
)


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := ""
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return replace_str
}

func StringJson_to_Json(str []uint8) map[string]interface{}{
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(str), &jsonMap)
	return jsonMap
}


func PrettyString(str string) (string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return ""
	}
	return prettyJSON.String()
}
