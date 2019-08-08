package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *BridgeOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.BridgeFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Bridge
	for _, res := range results {
		dest := &sacloud.Bridge{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.BridgeFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Bridges: values,
	}, nil
}

// Create is fake implementation
func (o *BridgeOp) Create(ctx context.Context, zone string, param *sacloud.BridgeCreateRequest) (*sacloud.Bridge, error) {
	result := &sacloud.Bridge{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	putBridge(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *BridgeOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Bridge, error) {
	value := getBridgeByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Bridge{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *BridgeOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.BridgeUpdateRequest) (*sacloud.Bridge, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	putBridge(zone, value)

	return value, nil
}

// Patch is fake implementation
func (o *BridgeOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.BridgePatchRequest) (*sacloud.Bridge, error) {
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

	putBridge(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *BridgeOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}
