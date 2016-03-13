package machine

import (
	"testing"

	machine "github.com/justicezyx/ResourceManager/machine"
	service_pb "github.com/justicezyx/ResourceManager/proto"
)

func TestSum(t *testing.T) {
	cases := []struct {
		in   service_pb.ReservationGroup
		exp  service_pb.ResourceUnit
		desc string
	}{
		{
			in: service_pb.ReservationGroup{
				ResourceUnit: &service_pb.ResourceUnit{
					Cpu:    1,
					Memory: 1,
					Disk:   1,
				},
				Machines: []string{"a", "b", "c"},
			},
			exp: service_pb.ResourceUnit{
				Cpu:    3,
				Memory: 3,
				Disk:   3,
			},
			desc: "Sum should get the result",
		},
		{
			in: service_pb.ReservationGroup{
				ResourceUnit: &service_pb.ResourceUnit{
					Cpu:    1,
					Memory: 1,
					Disk:   1,
				},
			},
			exp: service_pb.ResourceUnit{
				Cpu:    0,
				Memory: 0,
				Disk:   0,
			},
			desc: "Empty machines",
		},
		{
			in: service_pb.ReservationGroup{
				ResourceUnit: &service_pb.ResourceUnit{
					Cpu:    1,
					Memory: 1,
				},
				Machines: []string{"a", "b", "c"},
			},
			exp: service_pb.ResourceUnit{
				Cpu:    3,
				Memory: 3,
				Disk:   0,
			},
			desc: "Default value is 0",
		},
	}

	for _, c := range cases {
		if got := machine.Sum(c.in); got != c.exp {
			t.Errorf("%s, got: %v, expected: %v", c.desc, got, c.exp)
		}
	}
}
