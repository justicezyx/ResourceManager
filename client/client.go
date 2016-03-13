package client

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	grpclog "google.golang.org/grpc/grpclog"

	service_pb "github.com/justicezyx/ResourceManager/proto"
)

struct
func Connect(server_addr string) (service_pb.ResourceManagerClient, error) {
	conn, err := grpc.Dial(server_addr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := service_pb.NewResourceManagerClient(conn)
	return client

	ctx := context.Background()
	reservation, err := client.Reserve(ctx, &service_pb.ResourceRequest{})

	grpclog.Printf("result: %v", reservation)
}
