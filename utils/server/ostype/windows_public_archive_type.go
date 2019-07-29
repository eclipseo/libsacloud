package ostype

import "github.com/sacloud/libsacloud/v2/sacloud/ostype"

//go:generate stringer -type=WindowsPublicArchiveType

// WindowsPublicArchiveType Windows系パブリックアーカイブ
type WindowsPublicArchiveType int

const (
	// Windows2016 OS種別:Windows Server 2016 Datacenter Edition
	Windows2016 WindowsPublicArchiveType = iota
	// Windows2016RDS OS種別:Windows Server 2016 RDS
	Windows2016RDS
	// Windows2016RDSOffice OS種別:Windows Server 2016 RDS(Office)
	Windows2016RDSOffice
	// Windows2016SQLServerWeb OS種別:Windows Server 2016 SQLServer(Web)
	Windows2016SQLServerWeb
	// Windows2016SQLServerStandard OS種別:Windows Server 2016 SQLServer 2016(Standard)
	Windows2016SQLServerStandard
	// Windows2016SQLServer2017Standard OS種別:Windows Server 2016 SQLServer 2017(Standard)
	Windows2016SQLServer2017Standard
	// Windows2016SQLServerStandardAll OS種別:Windows Server 2016 SQLServer(Standard) + RDS + Office
	Windows2016SQLServerStandardAll
	// Windows2016SQLServer2017StandardAll OS種別:Windows Server 2016 SQLServer 2017(Standard) + RDS + Office
	Windows2016SQLServer2017StandardAll
	// Windows2019 OS種別:Windows Server 2019 Datacenter Edition
	Windows2019
)

var WindowsPublicArchives = map[WindowsPublicArchiveType]ostype.ArchiveOSType{
	Windows2016:                         ostype.Windows2016,
	Windows2016RDS:                      ostype.Windows2016RDS,
	Windows2016RDSOffice:                ostype.Windows2016RDSOffice,
	Windows2016SQLServerWeb:             ostype.Windows2016SQLServerWeb,
	Windows2016SQLServerStandard:        ostype.Windows2016SQLServerStandard,
	Windows2016SQLServer2017Standard:    ostype.Windows2016SQLServer2017Standard,
	Windows2016SQLServerStandardAll:     ostype.Windows2016SQLServerStandardAll,
	Windows2016SQLServer2017StandardAll: ostype.Windows2016SQLServer2017StandardAll,
	Windows2019:                         ostype.Windows2019,
}
