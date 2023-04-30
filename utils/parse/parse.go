package parse

import "encoding/json"


func StrToStruct(text string, strt  interface{}) error {

	err := json.Unmarshal([]byte(text), &strt)

	return err
}