gen:
	protoc -I=proto --go_out=. --go-grpc_out=. proto/*.proto
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
