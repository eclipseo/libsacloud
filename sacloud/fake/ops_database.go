package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *DatabaseOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.DatabaseFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Database
	for _, res := range results {
		dest := &sacloud.Database{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.DatabaseFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		Databases: values,
	}, nil
}

// Create is fake implementation
func (o *DatabaseOp) Create(ctx context.Context, zone string, param *sacloud.DatabaseCreateRequest) (*sacloud.Database, error) {
	result := &sacloud.Database{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "database"
	result.Availability = types.Availabilities.Available

	putDatabase(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})
	return result, nil
}

// Read is fake implementation
func (o *DatabaseOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Database, error) {
	value := getDatabaseByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Database{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DatabaseOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.DatabaseUpdateRequest) (*sacloud.Database, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDatabase(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *DatabaseOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.DatabasePatchRequest) (*sacloud.Database, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, value); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	copySameNameField(param, value)

	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}
	if param.PatchEmptyToCommonSetting {
		value.CommonSetting = nil
	}
	if param.PatchEmptyToBackupSetting {
		value.BackupSetting = nil
	}
	if param.PatchEmptyToReplicationSetting {
		value.ReplicationSetting = nil
	}

	putDatabase(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *DatabaseOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *DatabaseOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// Boot is fake implementation
func (o *DatabaseOp) Boot(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// Shutdown is fake implementation
func (o *DatabaseOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// Reset is fake implementation
func (o *DatabaseOp) Reset(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return nil
}

// MonitorCPU is fake implementation
func (o *DatabaseOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.CPUTimeActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.CPUTimeActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorCPUTimeValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			CPUTime: float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorDisk is fake implementation
func (o *DatabaseOp) MonitorDisk(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.DiskActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.DiskActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorDiskValue{
			Time:  now.Add(time.Duration(i*-5) * time.Minute),
			Read:  float64(random(1000)),
			Write: float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorInterface is fake implementation
func (o *DatabaseOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.InterfaceActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.InterfaceActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorInterfaceValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			Receive: float64(random(1000)),
			Send:    float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorDatabase is fake implementation
func (o *DatabaseOp) MonitorDatabase(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.DatabaseActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.DatabaseActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorDatabaseValue{
			Time:              now.Add(time.Duration(i*-5) * time.Minute),
			TotalMemorySize:   float64(random(1000)),
			UsedMemorySize:    float64(random(1000)),
			TotalDisk1Size:    float64(random(1000)),
			UsedDisk1Size:     float64(random(1000)),
			TotalDisk2Size:    float64(random(1000)),
			UsedDisk2Size:     float64(random(1000)),
			BinlogUsedSizeKiB: float64(random(1000)),
			DelayTimeSec:      float64(random(1000)),
		})
	}

	return res, nil
}

// Status is fake implementation
func (o *DatabaseOp) Status(ctx context.Context, zone string, id types.ID) (*sacloud.DatabaseStatus, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &sacloud.DatabaseStatus{
		Status:  value.InstanceStatus,
		IsFatal: true,
		Version: &sacloud.DatabaseVersionInfo{
			LastModified: value.ModifiedAt.String(),
			CommitHash:   "foobar",
			Status:       "up",
			Tag:          "stable",
		},
	}, nil
}