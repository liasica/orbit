// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package app

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()
	return &Validator{validator: validate}
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}
