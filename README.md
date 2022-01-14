# basic_grpc
Basic tutorial grpc 

# Compilar protoc
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative records.proto
