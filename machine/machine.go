// Machine provides a daemon that does resource accounting and contacts ResourceManager
// for registration.
package machine

import service_pb "github.com/justicezyx/ResourceManager/proto"

func Sum(reservation_group service_pb.ReservationGroup) service_pb.ResourceUnit {
	num_of_unit := len(reservation_group.Machines)
	return service_pb.ResourceUnit{
		Cpu:    int32(num_of_unit) * reservation_group.ResourceUnit.Cpu,
		Memory: int32(num_of_unit) * reservation_group.ResourceUnit.Memory,
		Disk:   int64(num_of_unit) * reservation_group.ResourceUnit.Disk,
	}
}
