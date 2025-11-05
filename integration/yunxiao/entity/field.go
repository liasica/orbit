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
