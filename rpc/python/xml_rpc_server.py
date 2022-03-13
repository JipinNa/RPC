from xmlrpc.server import SimpleXMLRPCServer as Server

class Calculate:
    def add(self, x, y):
        return x + y
    def multiply(self, x, y):
        return x * y
    def subtract(self, x, y):
        return abs(x - y)
    def divide(self, x, y):
        if (y == 0):
            return 0
        return x / y
obj = Calculate()
srv = Server(("localhost", 8090))
srv.register_instance(obj)
print("start listenning on port 8090")
srv.serve_forever()

# if __name__ == "__main__":
#     obj = Calculate()
#     srv = Server(("localhost", 8088))
#     srv.register_instance(obj)
#     print("start listenning on port 8090")
#     srv.serve_forever