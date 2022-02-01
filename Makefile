gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=z_generated --go_out=z_generated
clean:
	rm -r proto/*

server-1:
	go run server/server.go -port 9000
client-1:
	go run client/client.go -address 0.0.0.0:9000