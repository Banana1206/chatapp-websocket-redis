# Test Websocket

run server:
```
go run .\main.go
```

Run client on browser: 
```
let socket = new WebSocket("ws://localhost:3001/ws")

socket.onmessage = function(event) {
    console.log("Received message from server: " + event.data);
};

socket.send("hello World")
```