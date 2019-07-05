// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-result'; DO NOT EDIT

package sacloud

import "github.com/sacloud/libsacloud/v2/sacloud/naked"

// ArchiveFindResult represents the Result of API
type ArchiveFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Archives []*naked.Archive
}

// ArchiveCreateResult represents the Result of API
type ArchiveCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Archive *naked.Archive
}

// ArchiveCreateBlankResult represents the Result of API
type ArchiveCreateBlankResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Archive   *naked.Archive
	FTPServer *naked.OpeningFTPServer
}

// ArchiveReadResult represents the Result of API
type ArchiveReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Archive *naked.Archive
}

// ArchiveUpdateResult represents the Result of API
type ArchiveUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Archive *naked.Archive
}

// ArchiveOpenFTPResult represents the Result of API
type ArchiveOpenFTPResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	FTPServer *naked.OpeningFTPServer
}

// BridgeFindResult represents the Result of API
type BridgeFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Bridges []*naked.Bridge
}

// BridgeCreateResult represents the Result of API
type BridgeCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Bridge *naked.Bridge
}

// BridgeReadResult represents the Result of API
type BridgeReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Bridge *naked.Bridge
}

// BridgeUpdateResult represents the Result of API
type BridgeUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Bridge *naked.Bridge
}

// CDROMFindResult represents the Result of API
type CDROMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	CDROMs []*naked.CDROM
}

// CDROMCreateResult represents the Result of API
type CDROMCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CDROM     *naked.CDROM
	FTPServer *naked.OpeningFTPServer
}

// CDROMReadResult represents the Result of API
type CDROMReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CDROM *naked.CDROM
}

// CDROMUpdateResult represents the Result of API
type CDROMUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CDROM *naked.CDROM
}

// CDROMOpenFTPResult represents the Result of API
type CDROMOpenFTPResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	FTPServer *naked.OpeningFTPServer
}

// DiskFindResult represents the Result of API
type DiskFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Disks []*naked.Disk
}

// DiskCreateResult represents the Result of API
type DiskCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskCreateDistantlyResult represents the Result of API
type DiskCreateDistantlyResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskCreateWithConfigResult represents the Result of API
type DiskCreateWithConfigResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskCreateWithConfigDistantlyResult represents the Result of API
type DiskCreateWithConfigDistantlyResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskInstallDistantFromResult represents the Result of API
type DiskInstallDistantFromResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskInstallResult represents the Result of API
type DiskInstallResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskReadResult represents the Result of API
type DiskReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskUpdateResult represents the Result of API
type DiskUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Disk *naked.Disk
}

// DiskMonitorResult represents the Result of API
type DiskMonitorResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// GSLBFindResult represents the Result of API
type GSLBFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	CommonServiceItems []*naked.GSLB
}

// GSLBCreateResult represents the Result of API
type GSLBCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.GSLB
}

// GSLBReadResult represents the Result of API
type GSLBReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.GSLB
}

// GSLBUpdateResult represents the Result of API
type GSLBUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.GSLB
}

// InterfaceFindResult represents the Result of API
type InterfaceFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Interfaces []*naked.Interface
}

// InterfaceCreateResult represents the Result of API
type InterfaceCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Interface *naked.Interface
}

// InterfaceReadResult represents the Result of API
type InterfaceReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Interface *naked.Interface
}

// InterfaceUpdateResult represents the Result of API
type InterfaceUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Interface *naked.Interface
}

// InterfaceMonitorResult represents the Result of API
type InterfaceMonitorResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// InternetFindResult represents the Result of API
type InternetFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Internets []*naked.Internet
}

// InternetCreateResult represents the Result of API
type InternetCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Internet *naked.Internet
}

// InternetReadResult represents the Result of API
type InternetReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Internet *naked.Internet
}

// InternetUpdateResult represents the Result of API
type InternetUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Internet *naked.Internet
}

// InternetUpdateBandWidthResult represents the Result of API
type InternetUpdateBandWidthResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Internet *naked.Internet
}

// InternetAddSubnetResult represents the Result of API
type InternetAddSubnetResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Subnet *naked.Subnet
}

// InternetUpdateSubnetResult represents the Result of API
type InternetUpdateSubnetResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Subnet *naked.Subnet
}

// InternetMonitorResult represents the Result of API
type InternetMonitorResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// LoadBalancerFindResult represents the Result of API
type LoadBalancerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Appliances []*naked.LoadBalancer
}

// LoadBalancerCreateResult represents the Result of API
type LoadBalancerCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.LoadBalancer
}

// LoadBalancerReadResult represents the Result of API
type LoadBalancerReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.LoadBalancer
}

// LoadBalancerUpdateResult represents the Result of API
type LoadBalancerUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.LoadBalancer
}

// LoadBalancerMonitorInterfaceResult represents the Result of API
type LoadBalancerMonitorInterfaceResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// LoadBalancerStatusResult represents the Result of API
type LoadBalancerStatusResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	LoadBalancer []*naked.LoadBalancerStatus
}

// NFSFindResult represents the Result of API
type NFSFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Appliances []*naked.NFS
}

// NFSCreateResult represents the Result of API
type NFSCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.NFS
}

// NFSReadResult represents the Result of API
type NFSReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.NFS
}

// NFSUpdateResult represents the Result of API
type NFSUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.NFS
}

// NFSMonitorFreeDiskSizeResult represents the Result of API
type NFSMonitorFreeDiskSizeResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// NFSMonitorInterfaceResult represents the Result of API
type NFSMonitorInterfaceResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// NoteFindResult represents the Result of API
type NoteFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Notes []*naked.Note
}

// NoteCreateResult represents the Result of API
type NoteCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Note *naked.Note
}

// NoteReadResult represents the Result of API
type NoteReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Note *naked.Note
}

// NoteUpdateResult represents the Result of API
type NoteUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Note *naked.Note
}

// PacketFilterFindResult represents the Result of API
type PacketFilterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	PacketFilters []*naked.PacketFilter
}

// PacketFilterCreateResult represents the Result of API
type PacketFilterCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	PacketFilter *naked.PacketFilter
}

// PacketFilterReadResult represents the Result of API
type PacketFilterReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	PacketFilter *naked.PacketFilter
}

// PacketFilterUpdateResult represents the Result of API
type PacketFilterUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	PacketFilter *naked.PacketFilter
}

// ServerFindResult represents the Result of API
type ServerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Servers []*naked.Server
}

// ServerCreateResult represents the Result of API
type ServerCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Server *naked.Server
}

// ServerReadResult represents the Result of API
type ServerReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Server *naked.Server
}

// ServerUpdateResult represents the Result of API
type ServerUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Server *naked.Server
}

// ServerChangePlanResult represents the Result of API
type ServerChangePlanResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Server *naked.Server
}

// ServerMonitorResult represents the Result of API
type ServerMonitorResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// SIMFindResult represents the Result of API
type SIMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	CommonServiceItems []*naked.SIM
}

// SIMCreateResult represents the Result of API
type SIMCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.SIM
}

// SIMReadResult represents the Result of API
type SIMReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.SIM
}

// SIMUpdateResult represents the Result of API
type SIMUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	CommonServiceItem *naked.SIM
}

// SIMLogsResult represents the Result of API
type SIMLogsResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Logs []*naked.SIMLog
}

// SIMGetNetworkOperatorResult represents the Result of API
type SIMGetNetworkOperatorResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	NetworkOperationConfigs []*naked.SIMNetworkOperatorConfig
}

// SIMMonitorSIMResult represents the Result of API
type SIMMonitorSIMResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// SwitchFindResult represents the Result of API
type SwitchFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Switches []*naked.Switch
}

// SwitchCreateResult represents the Result of API
type SwitchCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Switch *naked.Switch
}

// SwitchReadResult represents the Result of API
type SwitchReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Switch *naked.Switch
}

// SwitchUpdateResult represents the Result of API
type SwitchUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Switch *naked.Switch
}

// VPCRouterFindResult represents the Result of API
type VPCRouterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Appliances []*naked.VPCRouter
}

// VPCRouterCreateResult represents the Result of API
type VPCRouterCreateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.VPCRouter
}

// VPCRouterReadResult represents the Result of API
type VPCRouterReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.VPCRouter
}

// VPCRouterUpdateResult represents the Result of API
type VPCRouterUpdateResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Appliance *naked.VPCRouter
}

// VPCRouterMonitorInterfaceResult represents the Result of API
type VPCRouterMonitorInterfaceResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Data *naked.MonitorValues
}

// ZoneFindResult represents the Result of API
type ZoneFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Zones []*naked.Zone
}

// ZoneReadResult represents the Result of API
type ZoneReadResult struct {
	IsOk bool `json:"is_ok,omitempty"` // is_ok

	Zone *naked.Zone
}
