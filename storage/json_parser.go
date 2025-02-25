package storage

import (
	"encoding/json"
	"gopasskeeper/constants"
	"log"
	"strings"
)

func SerializePasswordDataToJson(data PasswordJson) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(constants.ErrInternalSerialization)
	}
	return string(jsonData)
}

func DeserializePasswordDataFromJson(jsonData string) (PasswordJson, error) {

	//FIXME By default, encoding/json allows duplicate fields, using the last occurrence
	var data PasswordJson
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	return data, err
}
