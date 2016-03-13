package resource_manager

import (
	testing "testing"

	proto "github.com/golang/protobuf/proto"
	service "github.com/justicezyx/ResourceManager/proto"
	server "github.com/justicezyx/ResourceRequest/client"
	client "github.com/justicezyx/ResourceRequest/server"
)

func TestClientServer(t *testing.T) {
	go server.StartServer(10000)
}
