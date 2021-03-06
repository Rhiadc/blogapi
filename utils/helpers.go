package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BodyParser(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetID(r *http.Request) uint64 {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	return id
}
