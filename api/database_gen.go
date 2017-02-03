package api

/************************************************
  generated by IDE. for [DatabaseAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *DatabaseAPI) Reset() *DatabaseAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *DatabaseAPI) Offset(offset int) *DatabaseAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *DatabaseAPI) Limit(limit int) *DatabaseAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *DatabaseAPI) Include(key string) *DatabaseAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *DatabaseAPI) Exclude(key string) *DatabaseAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定項目でのフィルター
func (api *DatabaseAPI) FilterBy(key string, value interface{}) *DatabaseAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *DatabaseAPI) FilterMultiBy(key string, value interface{}) *DatabaseAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *DatabaseAPI) WithNameLike(name string) *DatabaseAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *DatabaseAPI) WithTag(tag string) *DatabaseAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *DatabaseAPI) WithTags(tags []string) *DatabaseAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *DatabaseAPI) WithSizeGib(size int) *DatabaseAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *DatabaseAPI) WithSharedScope() *DatabaseAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *DatabaseAPI) WithUserScope() *DatabaseAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *DatabaseAPI) SortBy(key string, reverse bool) *DatabaseAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *DatabaseAPI) SortByName(reverse bool) *DatabaseAPI {
	api.sortByName(reverse)
	return api
}

// func (api *DatabaseAPI) SortBySize(reverse bool) *DatabaseAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *DatabaseAPI) New() *sacloud.Database {
// 	return &sacloud.Database{}
// }

// func (api *DatabaseAPI) Create(value *sacloud.Database) (*sacloud.Database, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *DatabaseAPI) Read(id string) (*sacloud.Database, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *DatabaseAPI) Update(id string, value *sacloud.Database) (*sacloud.Database, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *DatabaseAPI) Delete(id string) (*sacloud.Database, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *DatabaseAPI) setStateValue(setFunc func(*sacloud.Request)) *DatabaseAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *DatabaseAPI) request(f func(*sacloud.Response) error) (*sacloud.Database, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.Database, nil
//}
//
//func (api *DatabaseAPI) createRequest(value *sacloud.Database) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.Database = value
//	return req
//}
