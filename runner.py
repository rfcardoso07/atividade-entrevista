import sys
import requests
import json

titulo = sys.argv[1]

response = requests.post("http://localhost:3000", titulo)

parsedResponse = json.loads(response.text)

documents = parsedResponse["docs"]

print("Títulos encontrados:\n")

for i, document in enumerate(documents, start=1):
    print(f'{i})', document["title"])