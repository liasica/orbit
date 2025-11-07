// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package yunxiao

import (
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

const (
	HookHeaderSecret     = "X-Projex-Signature"
	HookHeaderStatusFrom = "X-Status-From"
	HookHeaderStatusTo   = "X-Status-To"
)

type WebhookHandler func(data *entity.WebhookStatusEvent)

func StatusWebhook(r *http.Request, w http.ResponseWriter, handle WebhookHandler) {
	defer func() {
		// 直接返回 200 OK
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}()

	// 校验 secret
	if r.Header.Get(HookHeaderSecret) != instance.webhookSecret {
		log.Warn().Msg("云效 StatusWebhook Secret 校验失败")
		return
	}

	// 获取 status from / to
	from := r.Header.Get(HookHeaderStatusFrom)
	to := r.Header.Get(HookHeaderStatusTo)
	if from == "" || to == "" {
		log.Warn().Msg("云效 StatusWebhook 缺少状态变更信息")
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("读取 云效 StatusWebhook 请求体失败")
		return
	}

	data := &entity.WebhookStatusEvent{
		From: yc.WorkflowStatus(from),
		To:   yc.WorkflowStatus(to),
	}

	var node ast.Node
	node, err = sonic.Get(b, "identifier")
	if err != nil {
		log.Error().Err(err).Msg("获取工作项 identifier 失败")
		return
	}

	var identifier string
	identifier, err = node.String()
	if err != nil {
		log.Error().Err(err).Msg("解析工作项 identifier 失败")
		return
	}

	var workitem entity.Workitem
	workitem, err = GetWorkitem(identifier)
	if err != nil {
		log.Error().Err(err).Msgf("获取工作项信息失败: %s", identifier)
		return
	}

	data.Workitem = &workitem

	go handle(data)
}
