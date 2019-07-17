// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-fake-op'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
)

// SwitchFactoryFuncToFake switches sacloud.xxxAPI's factory methods to use fake client
func SwitchFactoryFuncToFake() {
	sacloud.SetClientFactoryFunc(ResourceArchive, func(caller sacloud.APICaller) interface{} {
		return NewArchiveOp()
	})
	sacloud.SetClientFactoryFunc(ResourceAuthStatus, func(caller sacloud.APICaller) interface{} {
		return NewAuthStatusOp()
	})
	sacloud.SetClientFactoryFunc(ResourceAutoBackup, func(caller sacloud.APICaller) interface{} {
		return NewAutoBackupOp()
	})
	sacloud.SetClientFactoryFunc(ResourceBill, func(caller sacloud.APICaller) interface{} {
		return NewBillOp()
	})
	sacloud.SetClientFactoryFunc(ResourceBridge, func(caller sacloud.APICaller) interface{} {
		return NewBridgeOp()
	})
	sacloud.SetClientFactoryFunc(ResourceCDROM, func(caller sacloud.APICaller) interface{} {
		return NewCDROMOp()
	})
	sacloud.SetClientFactoryFunc(ResourceCoupon, func(caller sacloud.APICaller) interface{} {
		return NewCouponOp()
	})
	sacloud.SetClientFactoryFunc(ResourceDatabase, func(caller sacloud.APICaller) interface{} {
		return NewDatabaseOp()
	})
	sacloud.SetClientFactoryFunc(ResourceDisk, func(caller sacloud.APICaller) interface{} {
		return NewDiskOp()
	})
	sacloud.SetClientFactoryFunc(ResourceDiskPlan, func(caller sacloud.APICaller) interface{} {
		return NewDiskPlanOp()
	})
	sacloud.SetClientFactoryFunc(ResourceDNS, func(caller sacloud.APICaller) interface{} {
		return NewDNSOp()
	})
	sacloud.SetClientFactoryFunc(ResourceGSLB, func(caller sacloud.APICaller) interface{} {
		return NewGSLBOp()
	})
	sacloud.SetClientFactoryFunc(ResourceIcon, func(caller sacloud.APICaller) interface{} {
		return NewIconOp()
	})
	sacloud.SetClientFactoryFunc(ResourceInterface, func(caller sacloud.APICaller) interface{} {
		return NewInterfaceOp()
	})
	sacloud.SetClientFactoryFunc(ResourceInternet, func(caller sacloud.APICaller) interface{} {
		return NewInternetOp()
	})
	sacloud.SetClientFactoryFunc(ResourceInternetPlan, func(caller sacloud.APICaller) interface{} {
		return NewInternetPlanOp()
	})
	sacloud.SetClientFactoryFunc(ResourceIPAddress, func(caller sacloud.APICaller) interface{} {
		return NewIPAddressOp()
	})
	sacloud.SetClientFactoryFunc(ResourceIPv6Net, func(caller sacloud.APICaller) interface{} {
		return NewIPv6NetOp()
	})
	sacloud.SetClientFactoryFunc(ResourceIPv6Addr, func(caller sacloud.APICaller) interface{} {
		return NewIPv6AddrOp()
	})
	sacloud.SetClientFactoryFunc(ResourceLicense, func(caller sacloud.APICaller) interface{} {
		return NewLicenseOp()
	})
	sacloud.SetClientFactoryFunc(ResourceLicenseInfo, func(caller sacloud.APICaller) interface{} {
		return NewLicenseInfoOp()
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
	sacloud.SetClientFactoryFunc(ResourcePrivateHost, func(caller sacloud.APICaller) interface{} {
		return NewPrivateHostOp()
	})
	sacloud.SetClientFactoryFunc(ResourcePrivateHostPlan, func(caller sacloud.APICaller) interface{} {
		return NewPrivateHostPlanOp()
	})
	sacloud.SetClientFactoryFunc(ResourceProxyLB, func(caller sacloud.APICaller) interface{} {
		return NewProxyLBOp()
	})
	sacloud.SetClientFactoryFunc(ResourceRegion, func(caller sacloud.APICaller) interface{} {
		return NewRegionOp()
	})
	sacloud.SetClientFactoryFunc(ResourceServer, func(caller sacloud.APICaller) interface{} {
		return NewServerOp()
	})
	sacloud.SetClientFactoryFunc(ResourceServerPlan, func(caller sacloud.APICaller) interface{} {
		return NewServerPlanOp()
	})
	sacloud.SetClientFactoryFunc(ResourceServiceClass, func(caller sacloud.APICaller) interface{} {
		return NewServiceClassOp()
	})
	sacloud.SetClientFactoryFunc(ResourceSIM, func(caller sacloud.APICaller) interface{} {
		return NewSIMOp()
	})
	sacloud.SetClientFactoryFunc(ResourceSimpleMonitor, func(caller sacloud.APICaller) interface{} {
		return NewSimpleMonitorOp()
	})
	sacloud.SetClientFactoryFunc(ResourceSSHKey, func(caller sacloud.APICaller) interface{} {
		return NewSSHKeyOp()
	})
	sacloud.SetClientFactoryFunc(ResourceSwitch, func(caller sacloud.APICaller) interface{} {
		return NewSwitchOp()
	})
	sacloud.SetClientFactoryFunc(ResourceVPCRouter, func(caller sacloud.APICaller) interface{} {
		return NewVPCRouterOp()
	})
	sacloud.SetClientFactoryFunc(ResourceWebAccel, func(caller sacloud.APICaller) interface{} {
		return NewWebAccelOp()
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
* AuthStatusOp
*************************************************/

// AuthStatusOp is fake implementation of AuthStatusAPI interface
type AuthStatusOp struct {
	key string
}

// NewAuthStatusOp creates new AuthStatusOp instance
func NewAuthStatusOp() sacloud.AuthStatusAPI {
	return &AuthStatusOp{
		key: ResourceAuthStatus,
	}
}

/*************************************************
* AutoBackupOp
*************************************************/

// AutoBackupOp is fake implementation of AutoBackupAPI interface
type AutoBackupOp struct {
	key string
}

// NewAutoBackupOp creates new AutoBackupOp instance
func NewAutoBackupOp() sacloud.AutoBackupAPI {
	return &AutoBackupOp{
		key: ResourceAutoBackup,
	}
}

/*************************************************
* BillOp
*************************************************/

// BillOp is fake implementation of BillAPI interface
type BillOp struct {
	key string
}

// NewBillOp creates new BillOp instance
func NewBillOp() sacloud.BillAPI {
	return &BillOp{
		key: ResourceBill,
	}
}

/*************************************************
* BridgeOp
*************************************************/

// BridgeOp is fake implementation of BridgeAPI interface
type BridgeOp struct {
	key string
}

// NewBridgeOp creates new BridgeOp instance
func NewBridgeOp() sacloud.BridgeAPI {
	return &BridgeOp{
		key: ResourceBridge,
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
* CouponOp
*************************************************/

// CouponOp is fake implementation of CouponAPI interface
type CouponOp struct {
	key string
}

// NewCouponOp creates new CouponOp instance
func NewCouponOp() sacloud.CouponAPI {
	return &CouponOp{
		key: ResourceCoupon,
	}
}

/*************************************************
* DatabaseOp
*************************************************/

// DatabaseOp is fake implementation of DatabaseAPI interface
type DatabaseOp struct {
	key string
}

// NewDatabaseOp creates new DatabaseOp instance
func NewDatabaseOp() sacloud.DatabaseAPI {
	return &DatabaseOp{
		key: ResourceDatabase,
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
* DiskPlanOp
*************************************************/

// DiskPlanOp is fake implementation of DiskPlanAPI interface
type DiskPlanOp struct {
	key string
}

// NewDiskPlanOp creates new DiskPlanOp instance
func NewDiskPlanOp() sacloud.DiskPlanAPI {
	return &DiskPlanOp{
		key: ResourceDiskPlan,
	}
}

/*************************************************
* DNSOp
*************************************************/

// DNSOp is fake implementation of DNSAPI interface
type DNSOp struct {
	key string
}

// NewDNSOp creates new DNSOp instance
func NewDNSOp() sacloud.DNSAPI {
	return &DNSOp{
		key: ResourceDNS,
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
* IconOp
*************************************************/

// IconOp is fake implementation of IconAPI interface
type IconOp struct {
	key string
}

// NewIconOp creates new IconOp instance
func NewIconOp() sacloud.IconAPI {
	return &IconOp{
		key: ResourceIcon,
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
* InternetOp
*************************************************/

// InternetOp is fake implementation of InternetAPI interface
type InternetOp struct {
	key string
}

// NewInternetOp creates new InternetOp instance
func NewInternetOp() sacloud.InternetAPI {
	return &InternetOp{
		key: ResourceInternet,
	}
}

/*************************************************
* InternetPlanOp
*************************************************/

// InternetPlanOp is fake implementation of InternetPlanAPI interface
type InternetPlanOp struct {
	key string
}

// NewInternetPlanOp creates new InternetPlanOp instance
func NewInternetPlanOp() sacloud.InternetPlanAPI {
	return &InternetPlanOp{
		key: ResourceInternetPlan,
	}
}

/*************************************************
* IPAddressOp
*************************************************/

// IPAddressOp is fake implementation of IPAddressAPI interface
type IPAddressOp struct {
	key string
}

// NewIPAddressOp creates new IPAddressOp instance
func NewIPAddressOp() sacloud.IPAddressAPI {
	return &IPAddressOp{
		key: ResourceIPAddress,
	}
}

/*************************************************
* IPv6NetOp
*************************************************/

// IPv6NetOp is fake implementation of IPv6NetAPI interface
type IPv6NetOp struct {
	key string
}

// NewIPv6NetOp creates new IPv6NetOp instance
func NewIPv6NetOp() sacloud.IPv6NetAPI {
	return &IPv6NetOp{
		key: ResourceIPv6Net,
	}
}

/*************************************************
* IPv6AddrOp
*************************************************/

// IPv6AddrOp is fake implementation of IPv6AddrAPI interface
type IPv6AddrOp struct {
	key string
}

// NewIPv6AddrOp creates new IPv6AddrOp instance
func NewIPv6AddrOp() sacloud.IPv6AddrAPI {
	return &IPv6AddrOp{
		key: ResourceIPv6Addr,
	}
}

/*************************************************
* LicenseOp
*************************************************/

// LicenseOp is fake implementation of LicenseAPI interface
type LicenseOp struct {
	key string
}

// NewLicenseOp creates new LicenseOp instance
func NewLicenseOp() sacloud.LicenseAPI {
	return &LicenseOp{
		key: ResourceLicense,
	}
}

/*************************************************
* LicenseInfoOp
*************************************************/

// LicenseInfoOp is fake implementation of LicenseInfoAPI interface
type LicenseInfoOp struct {
	key string
}

// NewLicenseInfoOp creates new LicenseInfoOp instance
func NewLicenseInfoOp() sacloud.LicenseInfoAPI {
	return &LicenseInfoOp{
		key: ResourceLicenseInfo,
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
* PrivateHostOp
*************************************************/

// PrivateHostOp is fake implementation of PrivateHostAPI interface
type PrivateHostOp struct {
	key string
}

// NewPrivateHostOp creates new PrivateHostOp instance
func NewPrivateHostOp() sacloud.PrivateHostAPI {
	return &PrivateHostOp{
		key: ResourcePrivateHost,
	}
}

/*************************************************
* PrivateHostPlanOp
*************************************************/

// PrivateHostPlanOp is fake implementation of PrivateHostPlanAPI interface
type PrivateHostPlanOp struct {
	key string
}

// NewPrivateHostPlanOp creates new PrivateHostPlanOp instance
func NewPrivateHostPlanOp() sacloud.PrivateHostPlanAPI {
	return &PrivateHostPlanOp{
		key: ResourcePrivateHostPlan,
	}
}

/*************************************************
* ProxyLBOp
*************************************************/

// ProxyLBOp is fake implementation of ProxyLBAPI interface
type ProxyLBOp struct {
	key string
}

// NewProxyLBOp creates new ProxyLBOp instance
func NewProxyLBOp() sacloud.ProxyLBAPI {
	return &ProxyLBOp{
		key: ResourceProxyLB,
	}
}

/*************************************************
* RegionOp
*************************************************/

// RegionOp is fake implementation of RegionAPI interface
type RegionOp struct {
	key string
}

// NewRegionOp creates new RegionOp instance
func NewRegionOp() sacloud.RegionAPI {
	return &RegionOp{
		key: ResourceRegion,
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
* ServerPlanOp
*************************************************/

// ServerPlanOp is fake implementation of ServerPlanAPI interface
type ServerPlanOp struct {
	key string
}

// NewServerPlanOp creates new ServerPlanOp instance
func NewServerPlanOp() sacloud.ServerPlanAPI {
	return &ServerPlanOp{
		key: ResourceServerPlan,
	}
}

/*************************************************
* ServiceClassOp
*************************************************/

// ServiceClassOp is fake implementation of ServiceClassAPI interface
type ServiceClassOp struct {
	key string
}

// NewServiceClassOp creates new ServiceClassOp instance
func NewServiceClassOp() sacloud.ServiceClassAPI {
	return &ServiceClassOp{
		key: ResourceServiceClass,
	}
}

/*************************************************
* SIMOp
*************************************************/

// SIMOp is fake implementation of SIMAPI interface
type SIMOp struct {
	key string
}

// NewSIMOp creates new SIMOp instance
func NewSIMOp() sacloud.SIMAPI {
	return &SIMOp{
		key: ResourceSIM,
	}
}

/*************************************************
* SimpleMonitorOp
*************************************************/

// SimpleMonitorOp is fake implementation of SimpleMonitorAPI interface
type SimpleMonitorOp struct {
	key string
}

// NewSimpleMonitorOp creates new SimpleMonitorOp instance
func NewSimpleMonitorOp() sacloud.SimpleMonitorAPI {
	return &SimpleMonitorOp{
		key: ResourceSimpleMonitor,
	}
}

/*************************************************
* SSHKeyOp
*************************************************/

// SSHKeyOp is fake implementation of SSHKeyAPI interface
type SSHKeyOp struct {
	key string
}

// NewSSHKeyOp creates new SSHKeyOp instance
func NewSSHKeyOp() sacloud.SSHKeyAPI {
	return &SSHKeyOp{
		key: ResourceSSHKey,
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
* VPCRouterOp
*************************************************/

// VPCRouterOp is fake implementation of VPCRouterAPI interface
type VPCRouterOp struct {
	key string
}

// NewVPCRouterOp creates new VPCRouterOp instance
func NewVPCRouterOp() sacloud.VPCRouterAPI {
	return &VPCRouterOp{
		key: ResourceVPCRouter,
	}
}

/*************************************************
* WebAccelOp
*************************************************/

// WebAccelOp is fake implementation of WebAccelAPI interface
type WebAccelOp struct {
	key string
}

// NewWebAccelOp creates new WebAccelOp instance
func NewWebAccelOp() sacloud.WebAccelAPI {
	return &WebAccelOp{
		key: ResourceWebAccel,
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
