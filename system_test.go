package resource_manager

import (
	testing "testing"

	context "golang.org/x/net/context"

	client "github.com/justicezyx/ResourceManager/client"
	service_pb "github.com/justicezyx/ResourceManager/proto"
	server "github.com/justicezyx/ResourceManager/server"
)

func TestClientServer(t *testing.T) {
	go server.StartServer(10000)
	c, err := client.Connect("localhost:10000")
	if err != nil {
		t.Error(err)
	}
	defer c.Connection.Close()
	ctx := context.Background()
	reservation, err := c.Reserve(ctx, &service_pb.ResourceRequest{
		Uid: "test",
	})
	if err != nil {
		t.Error(err)
	}

	if reservation.Uid != "test" {
		t.Errorf("Got %v expect %v", reservation, service_pb.ResourceReservation{
			Uid: "test",
		})
	}
}
