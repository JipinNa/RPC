import requests

rsp = requests.get("http://localhost:8083/?a=2&b=3")
print(rsp.text)