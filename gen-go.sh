PREFIX=github.com/turing-era/turingera-proto/pb

protoc -Iproto --go_out=module=${PREFIX}:pb proto/*.proto
protoc -Iproto --go-grpc_out=module=${PREFIX}:pb proto/*.proto
protoc -Iproto --grpc-gateway_out=module=${PREFIX},grpc_api_configuration=gateway-rules.yaml:pb proto/*.proto
