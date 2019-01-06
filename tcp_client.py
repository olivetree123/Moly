from socket import socket, AF_INET, SOCK_STREAM


s = socket(AF_INET, SOCK_STREAM)
s.connect(('localhost', 4004))
n = s.send(b'Hello')
resp = s.recv(8192)
print("resp = ", resp)