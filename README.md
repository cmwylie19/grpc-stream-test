# gRPC Tests
_This repo contains a gRPC application for testing with Envoy. The gRPC Server streaming endpoint sends 1 task into the stream per second to test timeout and keepAlive configurations._   
![Solo.io](https://www.solo.io/wp-content/uploads/2019/09/logo-w.svg)

## Tests
- ✔ Unary RPC
- ✔️ Server-Streaming RPC
- ✔️ Stream Interceptor
- ✔️ Unary Interceptor

## Gloo Edge
**Deploy grpc-tests app**   
```
kubectl apply -f k8s/grpc-tests.yaml
```

**Create a virtual service in Gloo Edge**
```
kubectl apply -f k8s/vs-grpc-test.yaml
```

**Apply access logging configuration to the gateway/gateway-proxy for logs (Optional)**   
(I am testing on AWS, if you are using GKE, set `useProxyProto: false` on line 35)
```
kubectl apply -f k8s/gateway-gateway-proxy.yaml
```

**Enable Gloo Edge Function Discovery Service**
```
kubectl label upstream -n gloo-system default-grpc-tests-8080 discovery.solo.io/function_discovery=enabled

kubectl get upstream -n gloo-system default-grpc-tests-8080 -o yaml
```

**Enumerate the available grpc services**
```
grpcurl -plaintext $(glooctl proxy address --port http) list
```

**Describe gRPC Services**
```
grpcurl -plaintext $(glooctl proxy address --port http) describe todo.Todo
```

**Create a single task**
```
grpcurl -plaintext -d '{"name":"Solo Test"}' $(glooctl proxy address --port http) todo.Todo/CreateTask
```

**[Create multiple tasks](#Gloo-Edge-CreateTask)** (Do this to test grpc server streaming endpoint)  
Click link 👆 to see create tasks.

**GetTasks (Unary RPC)**
```
grpcurl -plaintext $(glooctl proxy address --port http) todo.Todo/GetTasks
```

**GetTasks (server streaming RPC)**
```
grpcurl -plaintext  $(glooctl proxy address --port http) todo.Todo/GetTasksStream  
```

**Get Access Logs**
```
kubectl logs -f deploy/gateway-proxy -n gloo-system | grep '^{' | jq
```

**Get app logs**
```
kubectl logs -f deploy/grpc-tests
```

##### Gloo Edge CreateTask
```
grpcurl -plaintext -d '{"name":"test-1"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-2"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-3"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-4"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-5"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-6"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-7"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-8"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-9"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-10"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-11"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-12"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-13"}' $(glooctl proxy address --port http) todo.Todo/CreateTask

grpcurl -plaintext -d '{"name":"test-14"}' $(glooctl proxy address --port http) todo.Todo/CreateTask
```
