package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ArchiveOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Archive
	for _, res := range results {
		dest := &sacloud.Archive{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ArchiveFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Archives: values,
	}, nil
}

// Create is fake implementation
func (o *ArchiveOp) Create(ctx context.Context, zone string, param *sacloud.ArchiveCreateRequest) (*sacloud.Archive, error) {
	result := &sacloud.Archive{}

	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	if !param.SourceArchiveID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		diskOp := NewDiskOp()
		source, err := diskOp.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Availability
	}

	result.DisplayOrder = int64(random(100))
	result.Availability = types.Availabilities.Migrating
	result.DiskPlanID = types.DiskPlans.HDD
	result.DiskPlanName = "標準プラン"
	result.DiskPlanStorageClass = "iscsi9999"

	putArchive(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return result, nil
}

// CreateBlank is fake implementation
func (o *ArchiveOp) CreateBlank(ctx context.Context, zone string, param *sacloud.ArchiveCreateBlankRequest) (*sacloud.Archive, *sacloud.FTPServer, error) {
	result := &sacloud.Archive{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	result.Availability = types.Availabilities.Uploading

	putArchive(zone, result)

	return result, &sacloud.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("archive%d", result.ID),
		Password:  "password-is-not-a-password",
	}, nil
}

// Read is fake implementation
func (o *ArchiveOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Archive, error) {
	value := getArchiveByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Archive{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ArchiveOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ArchiveUpdateRequest) (*sacloud.Archive, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Patch if fake implementation
func (o *ArchiveOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.ArchivePatchRequest) (*sacloud.Archive, error) {
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

	putArchive(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *ArchiveOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// OpenFTP is fake implementation
func (o *ArchiveOp) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *sacloud.OpenFTPRequest) (*sacloud.FTPServer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	value.SetAvailability(types.Availabilities.Uploading)
	putArchive(zone, value)

	return &sacloud.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("archive%d", id),
		Password:  "password-is-not-a-password",
	}, nil
}

// CloseFTP is fake implementation
func (o *ArchiveOp) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if !value.Availability.IsUploading() {
		value.SetAvailability(types.Availabilities.Available)
	}
	putArchive(zone, value)
	return nil
}
