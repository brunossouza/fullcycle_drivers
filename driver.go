package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//error check
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func loadDrives(file string) []byte {
	jsonFile, err := os.Open(file)
	checkError(err)

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	checkError(err)

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
