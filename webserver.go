package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	http.HandleFunc("/", rootHandler)
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
