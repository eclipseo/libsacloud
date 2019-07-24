package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	serverAPIName     = "Server"
	serverAPIPathName = "server"
)

var serverAPI = &dsl.Resource{
	Name:       serverAPIName,
	PathName:   serverAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(serverAPIName, serverNakedType, findParameter, serverView),

		// create
		ops.Create(serverAPIName, serverNakedType, serverCreateParam, serverView),

		// read
		ops.Read(serverAPIName, serverNakedType, serverView),

		// update
		ops.Update(serverAPIName, serverNakedType, serverUpdateParam, serverView),

		// delete
		ops.Delete(serverAPIName),

		// change plan
		{
			ResourceName:    serverAPIName,
			Name:            "ChangePlan",
			PathFormat:      dsl.IDAndSuffixPathFormat("plan"),
			Method:          http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(serverChangePlanParam),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("plan", serverChangePlanParam),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: serverAPIName,
				Type: meta.Static(naked.Server{}),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Singular),
					DestField:   names.ResourceFieldName(serverAPIName, dsl.PayloadForms.Singular),
					IsPlural:    false,
					Model:       serverView,
				},
			},
		},

		// insert cdrom
		{
			ResourceName: serverAPIName,
			Name:         "InsertCDROM",
			PathFormat:   dsl.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: meta.Static(naked.CDROM{}),
					Name: "CDROM",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "insertParam",
					Type: &dsl.Model{
						Name: "InsertCDROMRequest",
						Fields: []*dsl.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					MapConvTag: "CDROM",
				},
			},
		},

		// eject cdrom
		{
			ResourceName: serverAPIName,
			Name:         "EjectCDROM",
			PathFormat:   dsl.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodDelete,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: meta.Static(naked.CDROM{}),
					Name: "CDROM",
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "ejectParam",
					Type: &dsl.Model{
						Name: "EjectCDROMRequest",
						Fields: []*dsl.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					MapConvTag: "CDROM",
				},
			},
		},

		// power management(boot/shutdown/reset)
		ops.Boot(serverAPIName),
		ops.Shutdown(serverAPIName),
		ops.Reset(serverAPIName),

		// send key
		{
			ResourceName: serverAPIName,
			Name:         "SendKey",
			PathFormat:   dsl.IDAndSuffixPathFormat("keyboard"),
			Method:       http.MethodPut,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("keyboardParam", serverSendKeyParam),
			},
			RequestEnvelope: dsl.RequestEnvelopeFromModel(serverSendKeyParam),
		},

		// get vnc proxy
		{
			ResourceName: serverAPIName,
			Name:         "GetVNCProxy",
			PathFormat:   dsl.IDAndSuffixPathFormat("vnc/proxy"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "VNCProxyInfo",
					Type: meta.Static(naked.VNCProxyInfo{}),
				},
			),
			Results: dsl.Results{
				{
					SourceField: "VNCProxyInfo",
					DestField:   serverVNCProxyView.Name,
					IsPlural:    false,
					Model:       serverVNCProxyView,
				},
			},
		},

		// monitor
		ops.Monitor(serverAPIName, monitorParameter, monitors.cpuTimeModel()),
	},
}

var (
	serverNakedType = meta.Static(naked.Server{})

	serverView = &dsl.Model{
		Name:      serverAPIName,
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.HostName(),
			fields.InterfaceDriver(),
			// server plan
			fields.ServerPlanID(),
			fields.ServerPlanName(),
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			// zone
			fields.Zone(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceBeforeStatus(),
			fields.InstanceStatusChangedAt(),
			fields.InstanceWarnings(),
			fields.InstanceWarningsValue(),

			// disks
			{
				Name: "Disks",
				Type: &dsl.Model{
					Name:      diskModel.Name,
					Fields:    diskModel.Fields,
					NakedType: meta.Static(naked.Disk{}),
					IsArray:   true,
				},
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",recursive",
				},
			},
			fields.Interfaces(),

			fields.CDROMID(),

			fields.PrivateHostID(),
			fields.PrivateHostName(),

			fields.BundleInfo(),

			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	serverCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(serverAPIName),
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			// server plan
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			{
				Name: "ConnectedSwitches",
				Type: &dsl.Model{
					Name: "ConnectedSwitch",
					Fields: []*dsl.FieldDesc{
						fields.ID(),
						fields.Scope(),
					},
					IsArray:   true,
					NakedType: meta.Static(naked.ConnectedSwitch{}),
				},
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: "[]ConnectedSwitches,recursive",
				},
			},
			fields.InterfaceDriver(),
			fields.HostName(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "WaitDiskMigration",
				Type: meta.TypeFlag,
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}

	serverUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(serverAPIName),
		NakedType: serverNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	serverChangePlanParam = &dsl.Model{
		Name: "ServerChangePlanRequest",
		Fields: []*dsl.FieldDesc{
			fields.CPU(),
			fields.MemoryMB(),
			fields.Generation(),
			fields.ServerPlanCommitment(),
		},
		NakedType: meta.Static(naked.ServerPlan{}),
	}

	serverSendKeyParam = &dsl.Model{
		Name: "SendKeyRequest",
		Fields: []*dsl.FieldDesc{
			fields.Def("Key", meta.TypeString),
			fields.Def("Keys", meta.TypeStringSlice),
		},
	}

	serverVNCProxyView = &dsl.Model{
		Name:      "VNCProxyInfo",
		NakedType: meta.Static(naked.VNCProxyInfo{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Status", meta.TypeString),
			fields.Def("Host", meta.TypeString),
			fields.Def("IOServerHost", meta.TypeString),
			fields.Def("Port", meta.TypeStringNumber),
			fields.Def("Password", meta.TypeString),
			fields.Def("VNCFile", meta.TypeString),
		},
	}
)
