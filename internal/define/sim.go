package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.SIM{})

	sim := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.SIMICCID(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.SIMProviderClass(),
			fields.SIMICCID(),
			fields.SIMPassCode(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	assignIPParam := &schema.Model{
		Name: "SIMAssignIPRequest",
		Fields: []*schema.FieldDesc{
			{
				Name: "IP",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMAssignIPRequest{}),
	}
	imeiLockParam := &schema.Model{
		Name: "SIMIMEILockRequest",
		Fields: []*schema.FieldDesc{
			{
				Name: "IMEI",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMIMEILockRequest{}),
	}

	simLog := &schema.Model{
		Name: "SIMLog",
		Fields: []*schema.FieldDesc{
			fields.New("Date", meta.TypeTime),
			fields.New("SessionStatus", meta.TypeString),
			fields.New("ResourceID", meta.TypeString),
			fields.New("IMEI", meta.TypeString),
			fields.New("IMSI", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMLog{}),
	}
	nwOpConfig := &schema.Model{
		Name: "SIMNetworkOperatorConfig",
		Fields: []*schema.FieldDesc{
			fields.New("Allow", meta.TypeFlag),
			fields.New("CountryCode", meta.TypeString),
			fields.New("Name", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMNetworkOperatorConfig{}),
	}
	nwOpConfigs := &schema.Model{
		Name: "SIMNetworkOperatorConfigs",
		Fields: []*schema.FieldDesc{
			{
				Name: "NetworkOperatorConfigs",
				Type: &schema.Model{
					Name: "SIMNetworkOperatorConfig",
					Fields: []*schema.FieldDesc{
						fields.New("Allow", meta.TypeFlag),
						fields.New("CountryCode", meta.TypeString),
						fields.New("Name", meta.TypeString),
					},
					NakedType: meta.Static(naked.SIMNetworkOperatorConfig{}),
					IsArray:   true,
				},
			},
		},
		NakedType: meta.Static(naked.SIMNetworkOperatorConfigs{}),
	}

	Resources.DefineWith("SIM", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationCommonServiceItemFind(nakedType, findParameter, sim),

			// create
			r.DefineOperationCommonServiceItemCreate(nakedType, createParam, sim),

			// read
			r.DefineOperationCommonServiceItemRead(nakedType, sim),

			// update
			r.DefineOperationCommonServiceItemUpdate(nakedType, updateParam, sim),

			// delete
			r.DefineOperationDelete(),

			// activate
			r.DefineSimpleOperation("Activate", http.MethodPut, "sim/activate"),
			// deactivate
			r.DefineSimpleOperation("Deactivate", http.MethodPut, "sim/deactivate"),

			// assignIP
			r.DefineOperation("AssignIP").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("sim/ip")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadName: "SIM",
					PayloadType: meta.Static(naked.SIMAssignIPRequest{}),
					Tags: &schema.FieldTags{
						JSON: "sim",
					},
				}).
				MappableArgument("param", assignIPParam),

			// clearIP
			r.DefineSimpleOperation("ClearIP", http.MethodDelete, "sim/ip"),

			// IMEILock
			r.DefineOperation("IMEILock").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("sim/imeilock")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadName: "SIM",
					PayloadType: meta.Static(naked.SIMIMEILockRequest{}),
					Tags: &schema.FieldTags{
						JSON: "sim",
					},
				}).
				MappableArgument("param", imeiLockParam),

			// IMEIUnlock
			r.DefineSimpleOperation("IMEIUnlock", http.MethodDelete, "sim/imeilock"),

			// Logs
			r.DefineOperation("Logs").
				Method(http.MethodGet).
				PathFormat(schema.IDAndSuffixPathFormat("sim/sessionlog")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				ResultPluralFromEnvelope(simLog, &schema.EnvelopePayloadDesc{
					PayloadName: "Logs",
					PayloadType: meta.Static(naked.SIMLog{}),
				}),

			// GetNetworkOperator
			r.DefineOperation("GetNetworkOperator").
				Method(http.MethodGet).
				PathFormat(schema.IDAndSuffixPathFormat("sim/network_operator_config")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				ResultPluralFromEnvelope(nwOpConfig, &schema.EnvelopePayloadDesc{
					PayloadName: "NetworkOperationConfigs",
					PayloadType: meta.Static(naked.SIMNetworkOperatorConfig{}),
				}),

			// SetNetworkOperator
			r.DefineOperation("SetNetworkOperator").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("sim/network_operator_config")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				PassthroughModelArgumentWithEnvelope("configs", nwOpConfigs),

			// monitor
			r.DefineOperationMonitorChild("SIM", "sim/metrics",
				monitorParameter, monitors.linkModel()),
		)
	}).PathName("commonserviceitemm")
}
