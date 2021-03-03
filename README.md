# greeter-api

gRPC api in Go for learning purposes, following [this](https://medium.com/unicoidtech/escrevendo-um-micro-servi%C3%A7o-grpc-em-go-estrutura-defini%C3%A7%C3%A3o-e-deploy-no-google-kubernetes-engine-82351e23a2f7) tutorial.


To test the project, run

```go run cmd/main.go```

```
grpcurl -import-path proto -proto greeting/v1/service.proto -plaintext -d @ localhost:5005 greeting.v1.GreeterService/Greet <<EOM
{"name":"Rafael"}
EOM
```

![](assets/example.png)
