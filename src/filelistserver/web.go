package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):]
			file, err := os.Open(path)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
			}
			defer file.Close()

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			writer.Write(bytes)
		})

	serve := http.ListenAndServe(":8888", nil)
	if serve != nil {
		panic(serve)
	}
}
