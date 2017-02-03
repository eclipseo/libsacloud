package api

/************************************************
  generated by IDE. for [CDROMAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件リセット
func (api *CDROMAPI) Reset() *CDROMAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *CDROMAPI) Offset(offset int) *CDROMAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *CDROMAPI) Limit(limit int) *CDROMAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *CDROMAPI) Include(key string) *CDROMAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *CDROMAPI) Exclude(key string) *CDROMAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *CDROMAPI) FilterBy(key string, value interface{}) *CDROMAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *CDROMAPI) FilterMultiBy(key string, value interface{}) *CDROMAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *CDROMAPI) WithNameLike(name string) *CDROMAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *CDROMAPI) WithTag(tag string) *CDROMAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *CDROMAPI) WithTags(tags []string) *CDROMAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// WithSizeGib サイズ条件
func (api *CDROMAPI) WithSizeGib(size int) *CDROMAPI {
	api.FilterBy("SizeMB", size*1024)
	return api
}

// WithSharedScope 公開スコープ条件
func (api *CDROMAPI) WithSharedScope() *CDROMAPI {
	api.FilterBy("Scope", "shared")
	return api
}

// WithUserScope ユーザースコープ条件
func (api *CDROMAPI) WithUserScope() *CDROMAPI {
	api.FilterBy("Scope", "user")
	return api
}

// SortBy 指定キーでのソート
func (api *CDROMAPI) SortBy(key string, reverse bool) *CDROMAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *CDROMAPI) SortByName(reverse bool) *CDROMAPI {
	api.sortByName(reverse)
	return api
}

// SortBySize サイズでのソート
func (api *CDROMAPI) SortBySize(reverse bool) *CDROMAPI {
	api.sortBy("SizeMB", reverse)
	return api
}

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// New 新規作成用パラメータ作成
func (api *CDROMAPI) New() *sacloud.CDROM {
	return &sacloud.CDROM{}
}

//func (api *CDROMAPI) Create(value *sacloud.CDROM) (*sacloud.CDROM, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.create(api.createRequest(value), res)
//	})
//}

// Read 読み取り
func (api *CDROMAPI) Read(id int64) (*sacloud.CDROM, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *CDROMAPI) Update(id int64, value *sacloud.CDROM) (*sacloud.CDROM, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *CDROMAPI) Delete(id int64) (*sacloud.CDROM, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *CDROMAPI) setStateValue(setFunc func(*sacloud.Request)) *CDROMAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *CDROMAPI) request(f func(*sacloud.Response) error) (*sacloud.CDROM, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.CDROM, nil
}

func (api *CDROMAPI) createRequest(value *sacloud.CDROM) *sacloud.Request {
	req := &sacloud.Request{}
	req.CDROM = value
	return req
}
