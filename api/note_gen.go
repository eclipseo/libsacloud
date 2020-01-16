// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

/************************************************
  generated by IDE. for [NoteAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *NoteAPI) Reset() *NoteAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *NoteAPI) Offset(offset int) *NoteAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *NoteAPI) Limit(limit int) *NoteAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *NoteAPI) Include(key string) *NoteAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *NoteAPI) Exclude(key string) *NoteAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *NoteAPI) FilterBy(key string, value interface{}) *NoteAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *NoteAPI) FilterMultiBy(key string, value interface{}) *NoteAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *NoteAPI) WithNameLike(name string) *NoteAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *NoteAPI) WithTag(tag string) *NoteAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *NoteAPI) WithTags(tags []string) *NoteAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *NoteAPI) WithSizeGib(size int) *NoteAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// WithSharedScope 公開スコープ条件
func (api *NoteAPI) WithSharedScope() *NoteAPI {
	api.FilterBy("Scope", "shared")
	return api
}

// WithUserScope ユーザースコープ条件
func (api *NoteAPI) WithUserScope() *NoteAPI {
	api.FilterBy("Scope", "user")
	return api
}

// SortBy 指定キーでのソート
func (api *NoteAPI) SortBy(key string, reverse bool) *NoteAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *NoteAPI) SortByName(reverse bool) *NoteAPI {
	api.sortByName(reverse)
	return api
}

// func (api *NoteAPI) SortBySize(reverse bool) *NoteAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *NoteAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *NoteAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *NoteAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *NoteAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *NoteAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *NoteAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *NoteAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *NoteAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// SetTag タグ条件
func (api *NoteAPI) SetTag(tag string) {
	api.FilterBy("Tags.Name", tag)
}

// SetTags タグ(複数)条件
func (api *NoteAPI) SetTags(tags []string) {
	api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *NoteAPI) SetSizeGib(size int) {
// 	api.FilterBy("SizeMB", size*1024)
// }

// SetSharedScope 公開スコープ条件
func (api *NoteAPI) SetSharedScope() {
	api.FilterBy("Scope", "shared")
}

// SetUserScope ユーザースコープ条件
func (api *NoteAPI) SetUserScope() {
	api.FilterBy("Scope", "user")
}

// SetSortBy 指定キーでのソート
func (api *NoteAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *NoteAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// func (api *NoteAPI) SetSortBySize(reverse bool) {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// New 新規作成用パラメーター作成
func (api *NoteAPI) New() *sacloud.Note {
	return &sacloud.Note{}
}

// Create 新規作成
func (api *NoteAPI) Create(value *sacloud.Note) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

// Read 読み取り
func (api *NoteAPI) Read(id int64) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *NoteAPI) Update(id int64, value *sacloud.Note) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *NoteAPI) Delete(id int64) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *NoteAPI) setStateValue(setFunc func(*sacloud.Request)) *NoteAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *NoteAPI) request(f func(*sacloud.Response) error) (*sacloud.Note, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Note, nil
}

func (api *NoteAPI) createRequest(value *sacloud.Note) *sacloud.Request {
	req := &sacloud.Request{}
	req.Note = value
	return req
}
