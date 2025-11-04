// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package repository

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/configure"
)

type ConfigureRepository struct {
	orm *ent.ConfigureClient
}

func NewConfigure() *ConfigureRepository {
	return &ConfigureRepository{
		orm: ent.Database.Configure,
	}
}

// GetValue 获取配置值
func (r *ConfigureRepository) GetValue(key configure.Key) (data sonic.NoCopyRawMessage, err error) {
	var result *ent.Configure
	result, err = ent.Database.Configure.Query().Where(configure.KeyEQ(key)).First(context.Background())
	if err != nil {
		return
	}
	data = result.Data
	return
}

// SetValue 设置配置值
func (r *ConfigureRepository) SetValue(key configure.Key, data sonic.NoCopyRawMessage) (err error) {
	return r.orm.Create().
		SetKey(key).
		SetData(data).
		OnConflictColumns(configure.FieldKey).
		UpdateNewValues().
		Exec(context.Background())
}
