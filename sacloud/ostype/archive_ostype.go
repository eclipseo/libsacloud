// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ostype is define OS type of SakuraCloud public archive
package ostype

//go:generate stringer -type=ArchiveOSType

// ArchiveOSType パブリックアーカイブOS種別
type ArchiveOSType int

const (
	// CentOS OS種別:CentOS
	CentOS ArchiveOSType = iota
	// CentOS6 OS種別:CentOS6
	CentOS6
	// Ubuntu OS種別:Ubuntu
	Ubuntu
	// Debian OS種別:Debian
	Debian
	// CoreOS OS種別:CoreOS
	CoreOS
	// RancherOS OS種別:RancherOS
	RancherOS
	// K3OS OS種別: k3OS
	K3OS
	// Kusanagi OS種別:Kusanagi(CentOS)
	Kusanagi
	// SophosUTM OS種別:Sophos UTM
	SophosUTM
	// FreeBSD OS種別:FreeBSD
	FreeBSD
	// Netwiser OS種別: Netwiser Virtual Edition
	Netwiser
	// OPNsense OS種別: OPNsense
	OPNsense
	// Windows2016 OS種別:Windows Server 2016 Datacenter Edition
	Windows2016
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
	// Custom OS種別:カスタム
	Custom
)

// ArchiveOSTypes アーカイブ種別のリスト
var ArchiveOSTypes = []ArchiveOSType{
	CentOS,
	CentOS6,
	Ubuntu,
	Debian,
	CoreOS,
	RancherOS,
	K3OS,
	Kusanagi,
	SophosUTM,
	FreeBSD,
	Netwiser,
	OPNsense,
	Windows2016,
	Windows2016RDS,
	Windows2016RDSOffice,
	Windows2016SQLServerWeb,
	Windows2016SQLServerStandard,
	Windows2016SQLServer2017Standard,
	Windows2016SQLServerStandardAll,
	Windows2016SQLServer2017StandardAll,
	Windows2019,
}

// OSTypeShortNames OSTypeとして利用できる文字列のリスト
var OSTypeShortNames = []string{
	"centos", "centos6", "ubuntu", "debian", "coreos",
	"rancheros", "k3os", "kusanagi", "sophos-utm", "freebsd",
	"netwiser", "opnsense",
	"windows2016", "windows2016-rds", "windows2016-rds-office",
	"windows2016-sql-web", "windows2016-sql-standard", "windows2016-sql-standard-all",
	"windows2016-sql2017-standard", "windows2016-sql2017-standard-all",
	"windows2019",
}

// IsWindows Windowsか
func (o ArchiveOSType) IsWindows() bool {
	switch o {
	case Windows2016, Windows2016RDS, Windows2016RDSOffice,
		Windows2016SQLServerWeb, Windows2016SQLServerStandard, Windows2016SQLServerStandardAll,
		Windows2016SQLServer2017Standard, Windows2016SQLServer2017StandardAll,
		Windows2019:
		return true
	default:
		return false
	}
}

// IsSupportDiskEdit ディスクの修正機能をフルサポートしているか(Windowsは一部サポートのためfalseを返す)
func (o ArchiveOSType) IsSupportDiskEdit() bool {
	switch o {
	case CentOS, CentOS6, Ubuntu, Debian, CoreOS, RancherOS, K3OS, Kusanagi, FreeBSD:
		return true
	default:
		return false
	}
}

// StrToOSType 文字列からArchiveOSTypesへの変換
func StrToOSType(osType string) ArchiveOSType {
	switch osType {
	case "centos":
		return CentOS
	case "centos6":
		return CentOS6
	case "ubuntu":
		return Ubuntu
	case "debian":
		return Debian
	case "coreos":
		return CoreOS
	case "rancheros":
		return RancherOS
	case "k3os":
		return K3OS
	case "kusanagi":
		return Kusanagi
	case "sophos-utm":
		return SophosUTM
	case "freebsd":
		return FreeBSD
	case "netwiser":
		return Netwiser
	case "opnsense":
		return OPNsense
	case "windows2016":
		return Windows2016
	case "windows2016-rds":
		return Windows2016RDS
	case "windows2016-rds-office":
		return Windows2016RDSOffice
	case "windows2016-sql-web":
		return Windows2016SQLServerWeb
	case "windows2016-sql-standard":
		return Windows2016SQLServerStandard
	case "windows2016-sql2017-standard":
		return Windows2016SQLServer2017Standard
	case "windows2016-sql-standard-all":
		return Windows2016SQLServerStandardAll
	case "windows2016-sql2017-standard-all":
		return Windows2016SQLServer2017StandardAll
	case "windows2019":
		return Windows2019
	default:
		return Custom
	}
}
