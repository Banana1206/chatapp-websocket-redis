# chatapp-websocket-redis

## Initialize
```
go mod init chatapp-websocket-redis
```

## Install third-party libraries

```
go get github.com/gorilla/mux
go get github.com/gorilla/websocket
go get github.com/rs/cors
go get github.com/go-redis/redis/v8
```

`mux` is for an HTTP server, `websocket` for WebSocket, `cors` to resolve CORS issues, and `redis` to connect with Redis.

## Run Server
- Http server
```
go run main.go -server=http
```

- Websocket server
```
go run main.go -server=websocket
```