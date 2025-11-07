// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-07, by liasica

package entity

import "github.com/liasica/orbit/config/yc"

type WebhookStatusEvent struct {
	From, To yc.WorkflowStatus
	Workitem *Workitem
}

type WebhookStatusEventWorkitem struct {
	Identifier string `json:"identifier"`
	Creator    struct {
		Identifier  string `json:"identifier"`
		RealName    string `json:"realName"`
		DisplayName string `json:"displayName"`
		NickName    string `json:"nickName"`
		Avatar      string `json:"avatar"`
	} `json:"creator"`
	GmtModified  int64  `json:"gmtModified"`
	SerialNumber int    `json:"serialNumber"`
	SpaceType    string `json:"spaceType"`
	WorkitemType struct {
		SystemDefault      string `json:"systemDefault"`
		Identifier         string `json:"identifier"`
		IsDeleted          bool   `json:"isDeleted"`
		DisplayName        string `json:"displayName"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		NameEn             string `json:"nameEn"`
		CategoryIdentifier string `json:"categoryIdentifier"`
	} `json:"workitemType"`
	Subject  string `json:"subject"`
	Modifier struct {
		Identifier  string `json:"identifier"`
		RealName    string `json:"realName"`
		DisplayName string `json:"displayName"`
		NickName    string `json:"nickName"`
		Avatar      string `json:"avatar"`
	} `json:"modifier"`
	Description      string `json:"description"`
	SpaceIdentifier  string `json:"spaceIdentifier"`
	CustomFieldValue []struct {
		FieldClassName     string `json:"fieldClassName"`
		FieldFormat        string `json:"fieldFormat"`
		FieldIdentifier    string `json:"fieldIdentifier"`
		WorkitemIdentifier string `json:"workitemIdentifier"`
		ValueList          []struct {
			DisplayValue string `json:"displayValue"`
			Identifier   string `json:"identifier"`
			ValueEn      string `json:"valueEn"`
			Value        string `json:"value"`
		} `json:"valueList"`
		Value string `json:"value"`
	} `json:"customFieldValue"`
	GmtCreate  int64 `json:"gmtCreate"`
	AssignedTo struct {
		Identifier  string `json:"identifier"`
		RealName    string `json:"realName"`
		DisplayName string `json:"displayName"`
		NickName    string `json:"nickName"`
		Avatar      string `json:"avatar"`
	} `json:"assignedTo"`
	Space struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		CustomCode string `json:"customCode"`
	} `json:"space"`
	OrganizationIdentifier string `json:"organizationIdentifier"`
	Category               string `json:"category"`
	Status                 struct {
		Identifier    string `json:"identifier"`
		DisplayName   string `json:"displayName"`
		Name          string `json:"name"`
		WorkflowStage struct {
			Identifier  string `json:"identifier"`
			DisplayName string `json:"displayName"`
			Name        string `json:"name"`
			NameEn      string `json:"nameEn"`
		} `json:"workflowStage"`
		NameEn string `json:"nameEn"`
	} `json:"status"`
}
