# chatapp-websocket-redis

## Initialize
```
go mod init gochatapp
```

## Install third-party libraries

```
go get github.com/gorilla/mux
go get github.com/gorilla/websocket
go get github.com/joho/godotenv
go get github.com/rs/cors
go get github.com/go-redis/redis/v8
```

`mux` is for an HTTP server, `websocket` for WebSocket, `godotenv` for reading .env, `cors` to resolve CORS issues, and `redis` to connect with Redis.