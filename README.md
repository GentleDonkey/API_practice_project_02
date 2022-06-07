###Required package
####https://github.com/gorilla/mux to set the server
####https://github.com/stretchr/testify to mock the data




###How to use
init go mod
```
go mod init
```
run go file
```
go run .
```
As I don't know where does the data come, 
I built a client.go file as a repository, 
to protect the data.
In this client.go file,
I mocked some data for testing my codes.
You could replace functions in client.go,
to test it with your data.

Thank you