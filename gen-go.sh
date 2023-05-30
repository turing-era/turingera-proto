PREFIX=github.com/turing-era/turingera-proto/go

protoc -Iproto --go_out=module=${PREFIX}:go proto/*.proto
protoc -Iproto --go-grpc_out=module=${PREFIX}:go proto/*.proto
protoc -Iproto --grpc-gateway_out=module=${PREFIX},grpc_api_configuration=gateway-rules.yaml:go proto/*.proto
