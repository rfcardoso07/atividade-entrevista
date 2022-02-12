package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	servidor := http.NewServeMux()
	servidor.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			fmt.Println("Requisição do tipo POST recebida!")
			requestBody, erro := ioutil.ReadAll(request.Body)
			if erro != nil {
				log.Fatal(erro)
			}
			fmt.Printf("Conteúdo da requisição: %s\n", requestBody)
		}
	})

	n := negroni.Classic()
	n.UseHandler(servidor)
	http.ListenAndServe(":3000", n)
	fmt.Println(n)
}
