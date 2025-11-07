// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package service

import (
	"github.com/rs/zerolog/log"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/integration/gitlab"
)

type GitlabService struct {
}

func NewGitlab() *GitlabService {
	return &GitlabService{}
}

func (s *GitlabService) WebhookHandler(_ git.EventType, data any) {
	switch v := data.(type) {
	case *git.MergeEvent:
		s.doMergeEvent(v)
	}
}

func (s *GitlabService) doMergeEvent(data *git.MergeEvent) {
	state := data.ObjectAttributes.State
	source := data.ObjectAttributes.SourceBranch
	target := data.ObjectAttributes.TargetBranch

	log.Info().Msgf("收到 Gitlab Merge Event (%s): %s => %s", state, source, target)

	switch state {
	case gitlab.MergeStateMerged:
		NewCooperation().GitlabMerged(source, target)
	}
}
