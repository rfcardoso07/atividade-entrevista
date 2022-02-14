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

	if request.Method == "POST" { //caso o método da requisição recebida seja injeção de dados (POST)

		requestBody, requestError := ioutil.ReadAll(request.Body) //lê entrada

		if requestError != nil {
			log.Fatal(requestError)
		}

		requestBodyString := string(requestBody) //converte para string

		query := strings.ReplaceAll(requestBodyString, " ", "+") //formata para realizar busca na Open Library

		openLibraryURL := "http://openlibrary.org/search.json?title=" + query

		//cria e efetua requisição para API externa
		libraryRequest, _ := http.NewRequest(http.MethodGet, openLibraryURL, nil)
		response, libraryError := http.DefaultClient.Do(libraryRequest)

		if libraryError != nil {
			log.Fatal(libraryError)
		}

		defer response.Body.Close()

		//lê e converte para string a resposta fornecida pela API
		libraryBody, _ := ioutil.ReadAll(response.Body)
		libraryString := string(libraryBody)

		//devolve a string como resultado da requisição recebida
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(writer, libraryString)
	}
}

func main() {

	servidor := http.NewServeMux()
	servidor.HandleFunc("/", handleRequest) //define função de processamento de requisições
	negroni := negroni.Classic()
	negroni.UseHandler(servidor)
	http.ListenAndServe(":3000", negroni) //aguarda e processa requisições na porta 3000
}
