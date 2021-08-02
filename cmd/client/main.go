package main

import (
	"context"
	"flag"
	"github.com/its-dastan/grpcDemo/pb"
	"github.com/its-dastan/grpcDemo/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server")

	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	//laptop.Id = "aa5654cf-f842-44a2-972c-d281b9b6687"

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}


	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop", err)
		}
		return
	}
	log.Printf("created laptop with id: %s",res.Id)

}
