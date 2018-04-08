package protocol

//go:generate protoc -I ../../vendor/ -I ../../vendor/github.com/gogo/protobuf/protobuf -I . --gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:./todo todo.proto
