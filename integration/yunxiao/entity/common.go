// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package entity

type IDName struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Label represents a label entity.
type Label struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// Status represents the status of a workitem.
type Status struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	NameEn      string `json:"nameEn,omitempty"`
}

// CustomFieldValueValue represents a value in customFieldValues.values.
type CustomFieldValueValue struct {
	DisplayValue string `json:"displayValue,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
}

// CustomFieldValue represents a custom field value.
type CustomFieldValue struct {
	FieldFormat string                  `json:"fieldFormat,omitempty"`
	FieldID     string                  `json:"fieldId,omitempty"`
	FieldName   string                  `json:"fieldName,omitempty"`
	Values      []CustomFieldValueValue `json:"values,omitempty"`
}
