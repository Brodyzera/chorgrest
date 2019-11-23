package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var metricsDir *string

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getFlatTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Endpoint Hit: search")

	data, err := ioutil.ReadFile(*metricsDir)
	checkError(err)
	fmt.Fprint(w, string(data))
}

func handleRequests() {
	http.HandleFunc("/tree/flat", getFlatTree)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	metricsDir = flag.String("metricsdir", "", "Directory used for reading metrics_flat.json")
	flag.Parse()

	handleRequests()
}
