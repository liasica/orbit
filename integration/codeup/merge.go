// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package codeup

import "time"

type MergeEvent struct {
	User             MergeUser             `json:"user,omitempty"`
	Version          string                `json:"version,omitempty"`
	ObjectKind       string                `json:"object_kind,omitempty"`
	ObjectAttributes MergeObjectAttributes `json:"object_attributes,omitempty"`
	Repository       MergeRepository       `json:"repository,omitempty"`
}

type MergeUser struct {
	AliyunPk  string `json:"aliyun_pk,omitempty"`
	ExternUid string `json:"extern_uid,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Name      string `json:"name,omitempty"`
}

type MergeObjectAttributes struct {
	State           string          `json:"state,omitempty"`
	BizId           string          `json:"biz_id,omitempty"`
	TargetProjectId int             `json:"target_project_id,omitempty"`
	WorkInProgress  bool            `json:"work_in_progress,omitempty"`
	Source          MergeProject    `json:"source,omitempty"`
	SourceProjectId int             `json:"source_project_id,omitempty"`
	Title           string          `json:"title,omitempty"`
	Url             string          `json:"url,omitempty"`
	SourceBranch    string          `json:"source_branch,omitempty"`
	UpdatedAt       time.Time       `json:"updated_at,omitempty"`
	ProjectId       int             `json:"project_id,omitempty"`
	AuthorAliyunPk  string          `json:"author_aliyun_pk,omitempty"`
	Action          string          `json:"action,omitempty"`
	TargetCommitId  string          `json:"target_commit_id,omitempty"`
	IsUpdateByPush  bool            `json:"is_update_by_push,omitempty"`
	LastCommit      MergeLastCommit `json:"last_commit,omitempty"`
	LocalId         int             `json:"local_id,omitempty"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	Description     string          `json:"description,omitempty"`
	SourceType      string          `json:"source_type,omitempty"`
	MergeStatus     string          `json:"merge_status,omitempty"`
	AuthorId        int             `json:"author_id,omitempty"`
	TargetBranch    string          `json:"target_branch,omitempty"`
	Target          MergeProject    `json:"target,omitempty"`
}

type MergeProject struct {
	SshUrl          string `json:"ssh_url,omitempty"`
	VisibilityLevel int    `json:"visibility_level,omitempty"`
	HttpUrl         string `json:"http_url,omitempty"`
	WebUrl          string `json:"web_url,omitempty"`
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
}

type MergeLastCommit struct {
	Author    MergeCommitAuthor `json:"author,omitempty"`
	Id        string            `json:"id,omitempty"`
	Message   string            `json:"message,omitempty"`
	Url       string            `json:"url,omitempty"`
	Timestamp time.Time         `json:"timestamp,omitempty"`
}

type MergeCommitAuthor struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type MergeRepository struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	GitHttpUrl  string `json:"git_http_url,omitempty"`
	GitSshUrl   string `json:"git_ssh_url,omitempty"`
	Url         string `json:"url,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}
