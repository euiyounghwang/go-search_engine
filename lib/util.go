package lib

import "log"


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := ""
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return replace_str
}