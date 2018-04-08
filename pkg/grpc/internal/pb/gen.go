package pb

//go:generate protoc -I ../../../../vendor/ -I ../../../../vendor/github.com/gogo/protobuf/protobuf -I ../../../protocol -I .. --gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mtodo.proto=go.zenithar.org/todo/pkg/protocol/todo:$GOPATH/src ../project.proto
