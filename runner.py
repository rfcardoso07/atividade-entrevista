import requests

titulo = "the lord of the rings"
request = requests.post("http://localhost:3000", titulo)
print(request.text)