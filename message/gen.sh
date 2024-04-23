protoc --go_out=plugins=grpc:./proto  ./proto/msg.proto
protoc --go_out=plugins=grpc:./proto  ./proto/wfl.proto

protoc --go_out=plugins=grpc:./proto  ./proto/paperio.proto

protoc --go_out=plugins=grpc:./proto  ./proto/packet.proto
protoc --go_out=plugins=grpc:./proto  ./proto/hamster.proto


protoc --go_out=plugins=grpc:./proto  ./proto/simple.proto
