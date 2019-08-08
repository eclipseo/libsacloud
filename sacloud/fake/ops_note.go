package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *NoteOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.NoteFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Note
	for _, res := range results {
		dest := &sacloud.Note{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.NoteFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Notes: values,
	}, nil
}

// Create is fake implementation
func (o *NoteOp) Create(ctx context.Context, param *sacloud.NoteCreateRequest) (*sacloud.Note, error) {
	result := &sacloud.Note{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	putNote(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *NoteOp) Read(ctx context.Context, id types.ID) (*sacloud.Note, error) {
	value := getNoteByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Note{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *NoteOp) Update(ctx context.Context, id types.ID, param *sacloud.NoteUpdateRequest) (*sacloud.Note, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putNote(sacloud.APIDefaultZone, value)
	return value, nil
}

// Patch is API call
func (o *NoteOp) Patch(ctx context.Context, id types.ID, param *sacloud.NotePatchRequest) (*sacloud.Note, error) {
	value, err := o.Read(ctx, id)
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

	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}
	if param.PatchEmptyToClass {
		value.Class = ""
	}
	if param.PatchEmptyToContent {
		value.Content = ""
	}

	putNote(sacloud.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *NoteOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
