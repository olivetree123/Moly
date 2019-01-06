from socketserver import BaseRequestHandler, TCPServer

class EchoHandler(BaseRequestHandler):
    def handle(self):
        while True:
            msg = self.request.recv(8192)
            if not msg:
                break
            print('Got connection from', self.client_address, "content = ", msg)
            self.request.send(b"HTTP/1.1 200 OK\nDate: Sat, 31 Dec 2005 23:59:59 GMT\nContent-Type: text/html;charset=ISO-8859-1\nContent-Length: 3\n\n123")

if __name__ == '__main__':
    serv = TCPServer(('', 4004), EchoHandler)
    serv.serve_forever()