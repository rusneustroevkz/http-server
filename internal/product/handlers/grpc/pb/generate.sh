rm -rf *.pb.go
protoc --go_out=. --go-grpc_out=. products.proto