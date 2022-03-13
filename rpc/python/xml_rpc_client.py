from xmlrpc import client

rsp = client.ServerProxy("http://localhost:8090")

print(rsp.add(2, 3))
print(rsp.multiply(2, 3))
print(rsp.subtract(2, 3))
print(rsp.divide(2, 3))