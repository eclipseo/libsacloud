package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *SwitchOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.SwitchFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Switch
	for _, res := range results {
		dest := &sacloud.Switch{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.SwitchFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Switches: values,
	}, nil
}

// Create is fake implementation
func (o *SwitchOp) Create(ctx context.Context, zone string, param *sacloud.SwitchCreateRequest) (*sacloud.Switch, error) {
	result := &sacloud.Switch{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Scope = types.Scopes.User
	putSwitch(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *SwitchOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Switch, error) {
	value := getSwitchByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Switch{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SwitchOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.SwitchUpdateRequest) (*sacloud.Switch, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putSwitch(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *SwitchOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.SwitchPatchRequest) (*sacloud.Switch, error) {
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

	if param.PatchEmptyToNetworkMaskLen {
		value.NetworkMaskLen = 0
	}
	if param.PatchEmptyToDefaultRoute {
		value.DefaultRoute = ""
	}
	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}

	putSwitch(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SwitchOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// ConnectToBridge is fake implementation
func (o *SwitchOp) ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	bridgeOp := NewBridgeOp()
	bridge, err := bridgeOp.Read(ctx, zone, bridgeID)
	if err != nil {
		return fmt.Errorf("ConnectToBridge is failed: %s", err)
	}

	if bridge.SwitchInZone != nil {
		return newErrorConflict(o.key, id, fmt.Sprintf("Bridge[%d] already connected to switch", bridgeID))
	}

	value.BridgeID = bridgeID

	switchInZone := &sacloud.BridgeSwitchInfo{}
	copySameNameField(value, switchInZone)
	bridge.SwitchInZone = switchInZone

	//bridge.BridgeInfo = append(bridge.BridgeInfo, &sacloud.BridgeInfo{
	//	ID:     value.ID,
	//	Name:   value.Name,
	//	ZoneID: zoneIDs[zone],
	//})

	putBridge(zone, bridge)
	putSwitch(zone, value)
	return nil
}

// DisconnectFromBridge is fake implementation
func (o *SwitchOp) DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if value.BridgeID.IsEmpty() {
		return newErrorConflict(o.key, id, fmt.Sprintf("Switch[%d] already disconnected from switch", id))
	}

	bridgeOp := NewBridgeOp()
	bridge, err := bridgeOp.Read(ctx, zone, value.BridgeID)
	if err != nil {
		return fmt.Errorf("DisconnectFromBridge is failed: %s", err)
	}

	//var bridgeInfo []*sacloud.BridgeInfo
	//for _, i := range bridge.BridgeInfo {
	//	if i.ID != value.ID {
	//		bridgeInfo = append(bridgeInfo, i)
	//	}
	//}

	value.BridgeID = types.ID(0)
	bridge.SwitchInZone = nil
	// fakeドライバーではBridgeInfoに非対応
	//bridge.BridgeInfo = bridgeInfo

	putBridge(zone, bridge)
	putSwitch(zone, value)
	return nil
}

// GetServers is fake implementation
func (o *SwitchOp) GetServers(ctx context.Context, zone string, id types.ID) (*sacloud.SwitchGetServersResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	res := &sacloud.SwitchGetServersResult{}
	if value.ServerCount == 0 {
		return res, nil
	}

	searched, err := NewServerOp().Find(ctx, zone, nil)
	for _, server := range searched.Servers {
		for _, nic := range server.Interfaces {
			if nic.SwitchID == id {
				res.Servers = append(res.Servers, server)
				res.Count++
				res.Total++
				break
			}
		}
	}
	return res, nil
}
