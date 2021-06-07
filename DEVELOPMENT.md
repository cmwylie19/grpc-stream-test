## Dev Mode

**Generate Proto**
```
protoc api/todo.proto --go_out=. --go-grpc_out=.
```

**Run the app**
```
go run main.go
```

**Send request to server**
```
// Available gRPC Services
grpcurl -plaintext localhost:8080 list

// Describe gRPC service
grpcurl -plaintext localhost:8080 describe todo.Todo

// Unary RPC CreateTask
grpcurl -plaintext -d '{"name":"test"}' localhost:8080 todo.Todo/CreateTask

// Unary RPC GetTasks
grpcurl -plaintext localhost:8080 todo.Todo/GetTasks

// Unary RPC GetTasks
grpcurl -plaintext  localhost:8080 todo.Todo/GetTasks 

// server streaming RPC GetTasksStream
grpcurl -plaintext  localhost:8080 todo.Todo/GetTasksStream  
```


**Fill Tasks to test stream**
```
grpcurl -plaintext -d '{"name":"test-1"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-2"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-3"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-4"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-5"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-6"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-7"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-8"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-9"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-10"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-11"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-12"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-13"}' localhost:8080 todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-14"}' localhost:8080 todo.Todo/CreateTask
```


**Build and Push**
```
docker build -t docker.io/cmwylie19/grpc-tests .; docker push docker.io/cmwylie19/grpc-tests
```