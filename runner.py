import sys
import requests
import json

#define título como o argumento passado via linha de comando
titulo = sys.argv[1]

#envia o título em requisição do tipo POST para a API Go (está ouvindo na porta 3000 do host)
response = requests.post("http://host.docker.internal:3000", titulo)

#extrai, formata e imprime na tela os títulos de livro retornados pela API

parsedResponse = json.loads(response.text)
documents = parsedResponse["docs"]

print("Títulos encontrados:\n")

for i, document in enumerate(documents, start=1):
    print(f'{i})', document["title"])