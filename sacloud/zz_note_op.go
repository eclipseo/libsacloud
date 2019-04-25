// generated by 'github.com/sacloud/libsacloud/tools/internal/gen-api-op'; DO NOT EDIT

package sacloud

import (
	"context"
	"encoding/json"
)

// NoteOp implements NoteAPI interface
type NoteOp struct {
	// Client APICaller
	Client APICaller
	// PathSuffix is used when building URL
	PathSuffix string
	// PathName is used when building URL
	PathName string
}

// NewNoteOp creates new NoteOp instance
func NewNoteOp(client APICaller) *NoteOp {
	return &NoteOp{
		Client:     client,
		PathSuffix: "api/cloud/1.1",
		PathName:   "note",
	}
}

// Create is API call
func (o *NoteOp) Create(ctx context.Context, zone string, param *NoteCreateRequest) (*NoteCommonResponse, error) {
	url, err := buildURL("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}", map[string]interface{}{
		"rootURL":    SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName":   o.PathName,
		"zone":       zone,
		"param":      param,
	})
	if err != nil {
		return nil, err
	}

	var body interface{}

	{
		n, err := param.ToNaked()
		if err != nil {
			return nil, err
		}
		body = &NakedNoteCreateRequest{
			Note: n,
		}
	}

	data, err := o.Client.Do(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	nakedResponse := &NakedNoteCreateResponse{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return nil, err
	}

	response0 := &NoteCommonResponse{}
	if err := response0.ParseNaked(nakedResponse.Note); err != nil {
		return nil, err
	}
	return response0, nil
}

// Read is API call
func (o *NoteOp) Read(ctx context.Context, zone string, id int64) (*NoteCommonResponse, error) {
	url, err := buildURL("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/{{.id}}", map[string]interface{}{
		"rootURL":    SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName":   o.PathName,
		"zone":       zone,
		"id":         id,
	})
	if err != nil {
		return nil, err
	}

	var body interface{}

	data, err := o.Client.Do(ctx, "GET", url, body)
	if err != nil {
		return nil, err
	}

	nakedResponse := &NakedNoteReadResponse{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return nil, err
	}

	response0 := &NoteCommonResponse{}
	if err := response0.ParseNaked(nakedResponse.Note); err != nil {
		return nil, err
	}
	return response0, nil
}

// Update is API call
func (o *NoteOp) Update(ctx context.Context, zone string, id int64, param *NoteUpdateRequest) (*NoteCommonResponse, error) {
	url, err := buildURL("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/{{.id}}", map[string]interface{}{
		"rootURL":    SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName":   o.PathName,
		"zone":       zone,
		"id":         id,
		"param":      param,
	})
	if err != nil {
		return nil, err
	}

	var body interface{}

	{
		n, err := param.ToNaked()
		if err != nil {
			return nil, err
		}
		body = &NakedNoteUpdateRequest{
			Note: n,
		}
	}

	data, err := o.Client.Do(ctx, "PUT", url, body)
	if err != nil {
		return nil, err
	}

	nakedResponse := &NakedNoteUpdateResponse{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return nil, err
	}

	response0 := &NoteCommonResponse{}
	if err := response0.ParseNaked(nakedResponse.Note); err != nil {
		return nil, err
	}
	return response0, nil
}

// Delete is API call
func (o *NoteOp) Delete(ctx context.Context, zone string, id int64) error {
	url, err := buildURL("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/{{.id}}", map[string]interface{}{
		"rootURL":    SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName":   o.PathName,
		"zone":       zone,
		"id":         id,
	})
	if err != nil {
		return err
	}

	var body interface{}

	data, err := o.Client.Do(ctx, "DELETE", url, body)
	if err != nil {
		return err
	}

	nakedResponse := &NakedNoteDeleteResponse{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return err
	}

	return nil
}
