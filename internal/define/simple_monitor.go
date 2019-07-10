package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	simpleMonitorAPIName     = "SimpleMonitor"
	simpleMonitorAPIPathName = "commonserviceitem"
)

var simpleMonitorAPI = &dsl.Resource{
	Name:       simpleMonitorAPIName,
	PathName:   simpleMonitorAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(simpleMonitorAPIName, simpleMonitorNakedType, findParameter, simpleMonitorView),

		// create
		ops.CreateCommonServiceItem(simpleMonitorAPIName, simpleMonitorNakedType, simpleMonitorCreateParam, simpleMonitorView),

		// read
		ops.ReadCommonServiceItem(simpleMonitorAPIName, simpleMonitorNakedType, simpleMonitorView),

		// update
		ops.UpdateCommonServiceItem(simpleMonitorAPIName, simpleMonitorNakedType, simpleMonitorUpdateParam, simpleMonitorView),

		// delete
		ops.Delete(simpleMonitorAPIName),

		// momitor
		ops.MonitorChild(simpleMonitorAPIName, "ResponseTime", "/activity/responsetimesec",
			monitorParameter, monitors.responseTimeSecModel()),

		// health check
		ops.HealthStatus(simpleMonitorAPIName, simpleMonitorHealthStatusNakedType, simpleMonitorHealthStatus),
	},
}

var (
	simpleMonitorNakedType             = meta.Static(naked.SimpleMonitor{})
	simpleMonitorHealthStatusNakedType = meta.Static(naked.SimpleMonitorHealthCheckStatus{})

	simpleMonitorView = &dsl.Model{
		Name:      simpleMonitorAPIName,
		NakedType: simpleMonitorNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.SimpleMonitorProviderClass(),

			// status
			fields.SimpleMonitorTarget(),

			// settings
			fields.SettingsHash(),
			fields.SimpleMonitorDelayLoop(),
			fields.SimpleMonitorEnabled(),
			// settings - health check
			fields.SimpleMonitorHealthCheck(),

			// settings - email
			fields.SimpleMonitorNotifyEmailEnabled(),
			fields.SimpleMonitorNotifyEmailHTML(),

			// settings - slack
			fields.SimpleMonitorNotifySlackEnabled(),
			fields.SimpleMonitorSlackWebhooksURL(),
		},
	}

	simpleMonitorCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(simpleMonitorAPIName),
		NakedType: simpleMonitorNakedType,
		Fields: []*dsl.FieldDesc{
			// creation time only
			fields.SimpleMonitorProviderClass(),
			{
				Name: "Target",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					Validate: "required",
					MapConv:  "Name/Status.Target", // NameとStatus.Targetに同じ値を設定
				},
			},

			// settings
			fields.SimpleMonitorDelayLoop(),
			fields.SimpleMonitorEnabled(),
			// settings - health check
			fields.SimpleMonitorHealthCheck(),

			// settings - email
			fields.SimpleMonitorNotifyEmailEnabled(),
			fields.SimpleMonitorNotifyEmailHTML(),

			// settings - slack
			fields.SimpleMonitorNotifySlackEnabled(),
			fields.SimpleMonitorSlackWebhooksURL(),

			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	simpleMonitorUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(simpleMonitorAPIName),
		NakedType: simpleMonitorNakedType,
		Fields: []*dsl.FieldDesc{
			// settings
			fields.SimpleMonitorDelayLoop(),
			fields.SimpleMonitorEnabled(),
			// settings - health check
			fields.SimpleMonitorHealthCheck(),

			// settings - email
			fields.SimpleMonitorNotifyEmailEnabled(),
			fields.SimpleMonitorNotifyEmailHTML(),

			// settings - slack
			fields.SimpleMonitorNotifySlackEnabled(),
			fields.SimpleMonitorSlackWebhooksURL(),

			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	simpleMonitorHealthStatus = &dsl.Model{
		Name:      "SimpleMonitorHealthStatus",
		NakedType: simpleMonitorHealthStatusNakedType,
		Fields: []*dsl.FieldDesc{
			{
				Name: "LastCheckedAt",
				Type: meta.TypeTime,
			},
			{
				Name: "LastHealthChangedAt",
				Type: meta.TypeTime,
			},
			{
				Name: "Health",
				Type: meta.Static(types.ESimpleMonitorHealth("")),
			},
		},
	}
)
