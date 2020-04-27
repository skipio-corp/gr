package main

import (
	"fmt"
	"io"
	"net"

	"github.com/skipio-corp/gr/service"
	"google.golang.org/grpc"
)

type server struct {
}

type contactStream struct {
}

func (s *server) UpsertContacts(stream service.ContactService_UpsertContactsServer) error {
	resp := &service.UpsertContactsResponse{}
	for {
		contact, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(resp)
		}
		if err != nil {
			return err
		}
		fmt.Printf("Got contact %v\n", contact)
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2323")
	if err != nil {
		panic(err)
	}
	s := &server{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service.RegisterContactServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
}
