// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package entity

type Field struct {
	Name             string                `json:"name,omitempty"`
	Description      string                `json:"description,omitempty"`
	Type             string                `json:"type,omitempty"`
	Format           string                `json:"format,omitempty"`
	DefaultValue     string                `json:"defaultValue,omitempty"`
	Options          []FieldOption         `json:"options,omitempty"`
	CascadingOptions *FieldCascadingOption `json:"cascadingOptions,omitempty"`
	Required         bool                  `json:"required,omitempty"`
	ShowWhenCreate   bool                  `json:"showWhenCreate,omitempty"`
	Id               string                `json:"id,omitempty"`
}
type FieldCascadingOption struct {
	MustSelectLeaf bool          `json:"mustSelectLeaf,omitempty"`
	OptionsList    []FieldOption `json:"optionsList,omitempty"`
}

type FieldOption struct {
	Value        string `json:"value,omitempty"`
	ValueEn      string `json:"valueEn,omitempty"`
	DisplayValue string `json:"displayValue,omitempty"`
	Id           string `json:"id,omitempty"`
}

// CustomField 自定义字段
type CustomField struct {
	DefaultValue    string   `json:"defaultValue,omitempty"`
	Description     string   `json:"description,omitempty"`
	DisabledOptions []string `json:"disabledOptions,omitempty"` // 字段待选值，只有字段是全局字段且是列表类型时才有效
	Name            string   `json:"name,omitempty"`
	NameEn          string   `json:"nameEn,omitempty"`
	OperatorId      string   `json:"operatorId,omitempty"`
	Options         []string `json:"options,omitempty"`
}
