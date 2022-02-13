package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/urfave/negroni"
)

func handleRequest(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {

		requestBody, requestError := ioutil.ReadAll(request.Body)

		if requestError != nil {
			log.Fatal(requestError)
		}

		requestBodyString := string(requestBody)
		query := strings.ReplaceAll(requestBodyString, " ", "+")

		openLibraryURL := "http://openlibrary.org/search.json?title=" + query
		libraryRequest, _ := http.NewRequest(http.MethodGet, openLibraryURL, nil)
		response, libraryError := http.DefaultClient.Do(libraryRequest)

		if libraryError != nil {
			log.Fatal(libraryError)
		}

		defer response.Body.Close()

		libraryBody, _ := ioutil.ReadAll(response.Body)
		libraryString := string(libraryBody)

		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(writer, libraryString)
	}
}

func main() {

	servidor := http.NewServeMux()
	servidor.HandleFunc("/", handleRequest)
	negroni := negroni.Classic()
	negroni.UseHandler(servidor)
	http.ListenAndServe(":3000", negroni)
}
