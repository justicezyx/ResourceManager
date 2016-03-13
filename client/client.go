package client

import (
	grpc "google.golang.org/grpc"

	service_pb "github.com/justicezyx/ResourceManager/proto"
)

type ResourceManagerClientWithConn struct {
	service_pb.ResourceManagerClient
	Connection *grpc.ClientConn
}

func Connect(server_addr string) (*ResourceManagerClientWithConn, error) {
	conn, err := grpc.Dial(server_addr, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		return nil, err
	}

	client := service_pb.NewResourceManagerClient(conn)
	return &ResourceManagerClientWithConn{
		ResourceManagerClient: client,
		Connection:            conn,
	}, nil
}
