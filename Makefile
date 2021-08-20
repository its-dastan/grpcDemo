gen:
	protoc -I=proto --go_out=. --go-grpc_out=. proto/*.proto --grpc-gateway_out=. --swagger_out=:swagger
	# protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative *.proto

clean: 
	rm pb/*.go

run:
	go run server.go

install:
	go get -u \
		google.golang.org/grpc
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

server: 
	go run cmd/server/main.go -port 8080

rest: 
	go run cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080

server1: 
	go run cmd/server/main.go -port 50051

server2: 
	go run cmd/server/main.go -port 50052

server1-tls: 
	go run cmd/server/main.go -port 50051 -tls

server2-tls: 
	go run cmd/server/main.go -port 50052 -tls

client: 
	go run cmd/client/main.go -address 0.0.0.0:8080

client-tls: 
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

evans:
	evans -r -p 8080

cert: 
	cd cert; sh gen.sh; cd ..

.PHONY: gen clean server client test cert