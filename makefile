run:
	HOST_NAME=localhost PORT=3333 PROTOCOL=tcp go run main.go
client:
	go run cmd/client/main.go

gocode:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/pokeapi.proto

swiftcode:
	protoc proto/pokeapi.proto \
    --swift_out=. \
    --swiftgrpc_out=Client=true,Server=false:.