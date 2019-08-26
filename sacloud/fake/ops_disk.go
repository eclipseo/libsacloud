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
func (o *DiskOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.DiskFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Disk
	for _, res := range results {
		dest := &sacloud.Disk{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.DiskFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Disks: values,
	}, nil
}

// Create is fake implementation
func (o *DiskOp) Create(ctx context.Context, zone string, param *sacloud.DiskCreateRequest, distantFrom []types.ID) (*sacloud.Disk, error) {
	result := &sacloud.Disk{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillDiskPlan)

	if result.Connection == types.EDiskConnection("") {
		result.Connection = types.DiskConnections.VirtIO
	}
	if !param.SourceArchiveID.IsEmpty() {
		archiveOp := NewArchiveOp()
		source, err := archiveOp.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Availability
	}

	putDisk(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		disk, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return disk, nil
	})

	return result, nil
}

// Config is fake implementation
func (o *DiskOp) Config(ctx context.Context, zone string, id types.ID, edit *sacloud.DiskEditRequest) error {
	// TODO ディスクに接続されたサーバのIDを拾ってInterfaces[0].UserSubnet.DefaultRoute/UserIPAddressなども書き換えた方がいいかも?
	return nil
}

// CreateWithConfig is fake implementation
func (o *DiskOp) CreateWithConfig(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, editParam *sacloud.DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*sacloud.Disk, error) {
	// check
	if !createParam.ServerID.IsEmpty() {
		serverOp := NewServerOp()
		_, err := serverOp.Read(ctx, zone, createParam.ServerID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), fmt.Sprintf("Server %s is not found", createParam.ServerID))
		}
	}

	result, err := o.Create(ctx, zone, createParam, distantFrom)
	if err != nil {
		return nil, err
	}

	if !createParam.ServerID.IsEmpty() && bootAtAvailable {
		waiter := sacloud.WaiterForReady(func() (interface{}, error) {
			disk, err := o.Read(ctx, zone, result.ID)
			if err != nil {
				return nil, err
			}
			return disk, nil
		})
		res, err := waiter.WaitForState(ctx)
		if err != nil {
			return nil, err
		}
		result = res.(*sacloud.Disk)

		// boot server
		serverOp := NewServerOp()
		if err := serverOp.Boot(ctx, zone, createParam.ServerID); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// ToBlank is fake implementation
func (o *DiskOp) ToBlank(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	value.SourceArchiveID = types.ID(0)
	value.SourceArchiveAvailability = types.Availabilities.Unknown
	value.SourceDiskID = types.ID(0)
	value.SourceDiskAvailability = types.Availabilities.Unknown

	putDisk(zone, value)
	return nil
}

// ResizePartition is fake implementation
func (o *DiskOp) ResizePartition(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// ConnectToServer is fake implementation
func (o *DiskOp) ConnectToServer(ctx context.Context, zone string, id types.ID, serverID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	serverOp := NewServerOp()
	server, err := serverOp.Read(ctx, zone, serverID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", serverID))
	}

	for _, connected := range server.Disks {
		if connected.ID == value.ID {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is already connected to Server[%d]", id, serverID))
		}
	}

	// TODO とりあえず同時実行制御は考慮しない。更新対象リソースが増えるようであれば実装方法を考える

	connectedDisk := &sacloud.ServerConnectedDisk{}
	copySameNameField(value, connectedDisk)
	server.Disks = append(server.Disks, connectedDisk)
	putServer(zone, server)
	value.ServerID = serverID
	putDisk(zone, value)

	return nil
}

// DisconnectFromServer is fake implementation
func (o *DiskOp) DisconnectFromServer(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if value.ServerID.IsEmpty() {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is not connected to Server", id))
	}

	serverOp := NewServerOp()
	server, err := serverOp.Read(ctx, zone, value.ServerID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", value.ServerID))
	}

	var disks []*sacloud.ServerConnectedDisk
	for _, connected := range server.Disks {
		if connected.ID != value.ID {
			connectedDisk := &sacloud.ServerConnectedDisk{}
			copySameNameField(value, connectedDisk)
			server.Disks = append(server.Disks, connectedDisk)
			disks = append(disks, connected)
		}
	}
	if len(disks) == len(server.Disks) {
		return newInternalServerError(o.key, id, fmt.Sprintf("Disk[%d] is not found on server's connected disks", id))
	}

	server.Disks = disks
	putServer(zone, server)
	value.ServerID = types.ID(0)
	putDisk(zone, value)

	return nil
}

// Install is fake implementation
func (o *DiskOp) Install(ctx context.Context, zone string, id types.ID, installParam *sacloud.DiskInstallRequest, distantFrom []types.ID) (*sacloud.Disk, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	fill(value, fillDiskPlan)
	putDisk(zone, value)
	return value, nil
}

// Read is fake implementation
func (o *DiskOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Disk, error) {
	value := getDiskByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Disk{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DiskOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.DiskUpdateRequest) (*sacloud.Disk, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDisk(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *DiskOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.DiskPatchRequest) (*sacloud.Disk, error) {
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
	if param.PatchEmptyToConnection {
		value.Connection = types.EDiskConnection("")
	}

	putDisk(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *DiskOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// Monitor is fake implementation
func (o *DiskOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.DiskActivity, error) {
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
