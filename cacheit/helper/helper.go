package helper

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/sauravrana646/cacheit/model"
)

func ErrNotNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func JsonToStruct(filename string) []model.Thread {
	var threads []model.Thread

	byteData, err := ioutil.ReadFile(filename)
	ErrNotNil(err)

	json.Unmarshal(byteData, &threads)

	return threads
}

func StructToJson(v any, filename string) (string , error) {
	resultByte, err := json.Marshal(v)
	ErrNotNil(err)

	err = ioutil.WriteFile(filename, resultByte, fs.ModeAppend)


	return "Operation Success" , err
}
