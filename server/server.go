package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"golang.org/x/net/context"

	service_pb "github.com/justicezyx/ResourceManager/proto"
	"github.com/justicezyx/ResourceManager/server"
)

func NewServer() *resourceManagerServer {
	return new(resourceManagerServer)
}

func StartServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	service_pb.RegisterResourceManagerServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}

type resourceManagerServer struct {
}

func (s *resourceManagerServer) Reserve(ctx context.Context, request *service_pb.ResourceRequest) (*service_pb.ResourceReservation, error) {
	fmt.Printf("Processing request: %v\n", request)
	return &service_pb.ResourceReservation{
		Uid: request.Uid,
	}, nil
}
