// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud-v2/sacloud"
)

// SwitchFactoryFuncToFake switches sacloud.xxxAPI's factory methods to use fake client
func SwitchFactoryFuncToFake() {
	sacloud.SetClientFactoryFunc(ResourceArchive, func(caller sacloud.APICaller) interface{} {
		return NewArchiveOp()
	})
	sacloud.SetClientFactoryFunc(ResourceCDROM, func(caller sacloud.APICaller) interface{} {
		return NewCDROMOp()
	})
	sacloud.SetClientFactoryFunc(ResourceDisk, func(caller sacloud.APICaller) interface{} {
		return NewDiskOp()
	})
	sacloud.SetClientFactoryFunc(ResourceGSLB, func(caller sacloud.APICaller) interface{} {
		return NewGSLBOp()
	})
	sacloud.SetClientFactoryFunc(ResourceInterface, func(caller sacloud.APICaller) interface{} {
		return NewInterfaceOp()
	})
	sacloud.SetClientFactoryFunc(ResourceLoadBalancer, func(caller sacloud.APICaller) interface{} {
		return NewLoadBalancerOp()
	})
	sacloud.SetClientFactoryFunc(ResourceNFS, func(caller sacloud.APICaller) interface{} {
		return NewNFSOp()
	})
	sacloud.SetClientFactoryFunc(ResourceNote, func(caller sacloud.APICaller) interface{} {
		return NewNoteOp()
	})
	sacloud.SetClientFactoryFunc(ResourcePacketFilter, func(caller sacloud.APICaller) interface{} {
		return NewPacketFilterOp()
	})
	sacloud.SetClientFactoryFunc(ResourceServer, func(caller sacloud.APICaller) interface{} {
		return NewServerOp()
	})
	sacloud.SetClientFactoryFunc(ResourceSwitch, func(caller sacloud.APICaller) interface{} {
		return NewSwitchOp()
	})
	sacloud.SetClientFactoryFunc(ResourceZone, func(caller sacloud.APICaller) interface{} {
		return NewZoneOp()
	})
}

/*************************************************
* ArchiveOp
*************************************************/

// ArchiveOp is fake implementation of ArchiveAPI interface
type ArchiveOp struct {
	key string
}

// NewArchiveOp creates new ArchiveOp instance
func NewArchiveOp() sacloud.ArchiveAPI {
	return &ArchiveOp{
		key: ResourceArchive,
	}
}

/*************************************************
* CDROMOp
*************************************************/

// CDROMOp is fake implementation of CDROMAPI interface
type CDROMOp struct {
	key string
}

// NewCDROMOp creates new CDROMOp instance
func NewCDROMOp() sacloud.CDROMAPI {
	return &CDROMOp{
		key: ResourceCDROM,
	}
}

/*************************************************
* DiskOp
*************************************************/

// DiskOp is fake implementation of DiskAPI interface
type DiskOp struct {
	key string
}

// NewDiskOp creates new DiskOp instance
func NewDiskOp() sacloud.DiskAPI {
	return &DiskOp{
		key: ResourceDisk,
	}
}

/*************************************************
* GSLBOp
*************************************************/

// GSLBOp is fake implementation of GSLBAPI interface
type GSLBOp struct {
	key string
}

// NewGSLBOp creates new GSLBOp instance
func NewGSLBOp() sacloud.GSLBAPI {
	return &GSLBOp{
		key: ResourceGSLB,
	}
}

/*************************************************
* InterfaceOp
*************************************************/

// InterfaceOp is fake implementation of InterfaceAPI interface
type InterfaceOp struct {
	key string
}

// NewInterfaceOp creates new InterfaceOp instance
func NewInterfaceOp() sacloud.InterfaceAPI {
	return &InterfaceOp{
		key: ResourceInterface,
	}
}

/*************************************************
* LoadBalancerOp
*************************************************/

// LoadBalancerOp is fake implementation of LoadBalancerAPI interface
type LoadBalancerOp struct {
	key string
}

// NewLoadBalancerOp creates new LoadBalancerOp instance
func NewLoadBalancerOp() sacloud.LoadBalancerAPI {
	return &LoadBalancerOp{
		key: ResourceLoadBalancer,
	}
}

/*************************************************
* NFSOp
*************************************************/

// NFSOp is fake implementation of NFSAPI interface
type NFSOp struct {
	key string
}

// NewNFSOp creates new NFSOp instance
func NewNFSOp() sacloud.NFSAPI {
	return &NFSOp{
		key: ResourceNFS,
	}
}

/*************************************************
* NoteOp
*************************************************/

// NoteOp is fake implementation of NoteAPI interface
type NoteOp struct {
	key string
}

// NewNoteOp creates new NoteOp instance
func NewNoteOp() sacloud.NoteAPI {
	return &NoteOp{
		key: ResourceNote,
	}
}

/*************************************************
* PacketFilterOp
*************************************************/

// PacketFilterOp is fake implementation of PacketFilterAPI interface
type PacketFilterOp struct {
	key string
}

// NewPacketFilterOp creates new PacketFilterOp instance
func NewPacketFilterOp() sacloud.PacketFilterAPI {
	return &PacketFilterOp{
		key: ResourcePacketFilter,
	}
}

/*************************************************
* ServerOp
*************************************************/

// ServerOp is fake implementation of ServerAPI interface
type ServerOp struct {
	key string
}

// NewServerOp creates new ServerOp instance
func NewServerOp() sacloud.ServerAPI {
	return &ServerOp{
		key: ResourceServer,
	}
}

/*************************************************
* SwitchOp
*************************************************/

// SwitchOp is fake implementation of SwitchAPI interface
type SwitchOp struct {
	key string
}

// NewSwitchOp creates new SwitchOp instance
func NewSwitchOp() sacloud.SwitchAPI {
	return &SwitchOp{
		key: ResourceSwitch,
	}
}

/*************************************************
* ZoneOp
*************************************************/

// ZoneOp is fake implementation of ZoneAPI interface
type ZoneOp struct {
	key string
}

// NewZoneOp creates new ZoneOp instance
func NewZoneOp() sacloud.ZoneAPI {
	return &ZoneOp{
		key: ResourceZone,
	}
}
