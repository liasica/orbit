// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package pingcode

import (
	"github.com/cockroachdb/pebble"

	"github.com/liasica/orbit/db"
)

var AuthTokenKey = []byte("PINGCODE_AUTH_TOKEN")

// AuthTokenResponse 获取企业令牌响应
type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
}

// RequestAuthToken 获取企业令牌
// https://open.pingcode.com/#api-%E5%AE%A2%E6%88%B7%E7%AB%AF%E5%87%AD%E6%8D%AE
func RequestAuthToken() (res *AuthTokenResponse, err error) {
	res = new(AuthTokenResponse)

	_, err = instance.client.R().
		SetQueryParam("grant_type", "client_credentials").
		SetQueryParam("client_id", instance.config.ClientID).
		SetQueryParam("client_secret", instance.config.Secret).
		SetResult(res).
		Get("/v1/auth/token")

	return
}

// StoreAuthToken 存储企业令牌
func StoreAuthToken(data *AuthTokenResponse) {
	db.Get().Set(AuthTokenKey, []byte(data.AccessToken), pebble.Sync)
}

func GetAuthToken() (token string) {
	return
}
