package main

import (
	"flag"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	service_pb "github.com/justicezyx/ResourceManager/proto"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address host:port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := service_pb.NewResourceManagerClient(conn)

	ctx := context.Background()
	reservation, err := client.Reserve(ctx, &service_pb.ResourceRequest{})

	grpclog.Printf("result: %v", reservation)
}
