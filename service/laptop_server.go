package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/its-dastan/grpcDemo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid:%v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot generate a new laptop ID:%v", err)
		}
		laptop.Id = id.String()
	}

	//time.Sleep(6*time.Second)

	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	}

	err := server.Store.Save(laptop)

	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil

}


func (server *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter:= req.GetFilter()
	log.Printf("received a search laptop request with filter : %v", filter)

	err:= server.Store.Search(
		stream.Context(),
		filter,
		func (laptop *pb.Laptop) error {
			res:= &pb.SearchLaptopResponse{Laptop: laptop}
			
			
			err:= stream.Send(res)
			if err!= nil{
				return err
			}
			log.Printf("sent laptop with id: %s",laptop.GetId())
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Aborted,"unexpected error: %v", err)
	}
	return nil
}