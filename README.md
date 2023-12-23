# go-search_engine


## Go Env
```bash
xcode-select --install
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
brew install go
go env
go version
```


## Run Go
```bash
go mod init go-search_engine
go run ./sample/hello.go
go build ./sample/hello.go && ./sample/hello
```


## Run Docker with sample
```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware

go-search_engine git:(master) ✗ go run ./main.go                             

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:9080


curl http://localhost:9080/
curl http://localhost:9080/health

{"time":"2023-12-23T14:39:42.970473-06:00","id":"","remote_ip":"127.0.0.1","host":"localhost:9080","method":"GET","uri":"/","user_agent":"curl/7.78.0","status":200,"error":"","latency":1875,"latency_human":"1.875µs","bytes_in":0,"bytes_out":13}
{"time":"2023-12-23T14:40:07.332221-06:00","id":"","remote_ip":"127.0.0.1","host":"localhost:9080","method":"GET","uri":"/health","user_agent":"curl/7.78.0","status":200,"error":"","latency":70208,"latency_human":"70.208µs","bytes_in":0,"bytes_out":14}
```
