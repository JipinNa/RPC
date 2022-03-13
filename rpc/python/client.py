from protobuf import hello_pb2

req = hello_pb2.HelloRequest()
req.name = "bbbb"
str = req.SerializeToString()

print(str)