package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	service_pb "github.com/justicezyx/ResourceManager/proto"
	server "github.com/justicezyx/ResourceManager/server"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	service_pb.RegisterResourceManagerServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
