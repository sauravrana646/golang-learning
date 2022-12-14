package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sauravrana646/cacheit/helper"
	"github.com/sauravrana646/cacheit/model"
)

const jsonfile string = "new.json"

func setBasicHeaders(w http.ResponseWriter, value string) http.ResponseWriter {
	w.Header().Set("Content-Type", value)
	return w
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w = setBasicHeaders(w, "text/html")
	w.Write([]byte(`<h2>Welcome to Cache It</h2>`))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w = setBasicHeaders(w, "application/json")

	// var threads []model.Thread

	byteData, err := ioutil.ReadFile(jsonfile)
	helper.ErrNotNil(err)

	// json.Unmarshal(byteData, &threads)

	w.Write(byteData)

	// fmt.Println("Data stored in the Threads is \n", threads)
}

func GetRaw() {
	var threads []model.Thread

	byteData, err := ioutil.ReadFile("saves.json")
	helper.ErrNotNil(err)

	json.Unmarshal(byteData, &threads)

	fmt.Println(threads)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	w = setBasicHeaders(w, "application/json")

	var result model.Thread

	query := mux.Vars(r)

	// fmt.Println("Vars are ", query)

	threads := helper.JsonToStruct(jsonfile)

	// fmt.Println("Threads are ", threads)

	// for _, thread := range threads {
	// 	fmt.Println("Thread id is ", thread.Id)
	// }

	for _, thread := range threads {
		if thread.Id == query["id"] {
			result = thread
			break
		}
	}

	byteResult, err := json.Marshal(result)

	helper.ErrNotNil(err)

	w.Write(byteResult)

	// fmt.Println("Data stored in the Threads is \n", threads)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	
	var updatedThread model.Thread
	var isIdFound bool = false 

	w = setBasicHeaders(w, "application/json")

	v := mux.Vars(r)

	dataByte, err := ioutil.ReadAll(r.Body)
	helper.ErrNotNil(err)

	json.Unmarshal(dataByte, &updatedThread)
	// fmt.Println("Vars are ", query)

	threads := helper.JsonToStruct(jsonfile)

	for key , thread := range threads {
		if thread.Id == v["id"]{
				threads = append(threads[:key], threads[key+1:]... )
				isIdFound = true
				break
		}
	}
	
	if !isIdFound{
		// w.Write([]byte(`No matching ID Found`))
		http.Error(w, "No matching ID Found", http.StatusBadRequest)
		return
	}

	// v := mux.Vars
}

func AddOne(w http.ResponseWriter, r *http.Request) {

	var newthread model.Thread

	threads:= helper.JsonToStruct(jsonfile)


	dataByte, err := ioutil.ReadAll(r.Body)
	helper.ErrNotNil(err)

	json.Unmarshal(dataByte, &newthread)

	fmt.Println("Thread to be added is ", newthread)

	for _, thread := range threads {
		if thread.Id == newthread.Id {
			http.Error(w, "Item with this ID already exists", http.StatusUnprocessableEntity)
			return
		}
	}

	threads = append(threads, newthread)

	res, err := helper.StructToJson(threads, jsonfile)
	
	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
	w.Write([]byte(res))

}

func DeleteOne(w http.ResponseWriter, r *http.Request)  {
	w =  setBasicHeaders(w, "text/plain")
	isIdFound := false

	v := mux.Vars(r)

	threads := helper.JsonToStruct(jsonfile)

	for key , thread := range threads {
		if thread.Id == v["id"]{
			threads = append(threads[:key], threads[key+1:]... )
			isIdFound = true
			break
		}
	}

	if !isIdFound{
		// w.Write([]byte(`No matching ID Found`))
		http.Error(w, "No matching ID Found", http.StatusBadRequest)
		return
	}

	res , err := helper.StructToJson(threads,jsonfile)

	if err != nil {
		http.Error(w, err.Error() , http.StatusInternalServerError)
	}
	w.Write([]byte(res))
}
