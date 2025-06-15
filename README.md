## 参考

go 中級者への道

## 構築手順

```
cd go_myapi
go mod init github.com/marozex/go_myapi
touch main.go
go get -u github.com/gorilla/mux
```

## 操作手順

サーバー起動

```
 go run main.go
```

疎通確認

```
curl http://localhost:1234/hello

curl http://localhost:1234/hello -X GET -w '%{http_code}\n'

curl http://localhost:1234/hello -X POST -w '%{http_code}\n'
```
