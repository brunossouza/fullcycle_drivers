package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadDrives(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	return data
}

//ListDrivers func to handler requests and expose the drivers list
func ListDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrives("drivers.json")
	w.Write([]byte(drivers))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/drivers", ListDrivers)
	http.ListenAndServe(":8080", r)
}
