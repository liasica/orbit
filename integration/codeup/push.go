// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package codeup

import "time"

type PushEvent struct {
	UserEmail         string         `json:"user_email,omitempty"`
	UserName          string         `json:"user_name,omitempty"`
	Repository        PushRepository `json:"repository,omitempty"`
	Ref               string         `json:"ref,omitempty"`
	ProjectId         int            `json:"project_id,omitempty"`
	After             string         `json:"after,omitempty"`
	TotalCommitsCount int            `json:"total_commits_count,omitempty"`
	Before            string         `json:"before,omitempty"`
	UserExternUid     string         `json:"user_extern_uid,omitempty"`
	CheckoutSha       string         `json:"checkout_sha,omitempty"`
	ObjectKind        string         `json:"object_kind,omitempty"`
	UserId            int            `json:"user_id,omitempty"`
	Commits           []PushCommit   `json:"commits,omitempty"`
	AliyunPk          string         `json:"aliyun_pk,omitempty"`
}

type PushRepository struct {
	Url             string `json:"url,omitempty"`
	Homepage        string `json:"homepage,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	VisibilityLevel int    `json:"visibility_level,omitempty"`
	GitHttpUrl      string `json:"git_http_url,omitempty"`
	GitSshUrl       string `json:"git_ssh_url,omitempty"`
}

type PushCommit struct {
	Url       string           `json:"url,omitempty"`
	Timestamp time.Time        `json:"timestamp,omitempty"`
	Author    PushCommitAuthor `json:"author,omitempty"`
	Id        string           `json:"id,omitempty"`
	Message   string           `json:"message,omitempty"`
}

type PushCommitAuthor struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
