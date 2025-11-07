// Copyright (C) moneta. 2025-present.
//
// Created at 2025-07-28, by liasica

package utils

import (
	"errors"
	"regexp"

	"resty.dev/v3"
)

var (
	ErrIpObtainFailed = errors.New("ip地址获取失败")
	IPRegexRule       = regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)
)

// GetMyIP 获取当前公网IP地址
// https://lolicp.com/others/202405106.html
// https://www.ip.cn/api/index?ip&type=0
func GetMyIP() (ip string, err error) {
	fs := []func() (string, error){
		getMyIPQQ,
		getMyIP189,
		getMyIP12306,
		getMyIPiQiyi,
	}

	for _, f := range fs {
		ip, err = f()
		if err == nil && ip != "" {
			return
		}
	}

	return
}

// 检查IP地址和错误, 返回格式化的错误
func getIPError(ip string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", ErrIpObtainFailed
	}
	return ip, nil
}

// 从12306获取IP地址
func getMyIP12306() (string, error) {
	// 获取IP地址
	var myIP struct {
		IP string `json:"di"`
	}
	_, err := resty.New().R().SetResult(&myIP).Get(`https://exservice.12306.cn/excater/bonree/grip`)
	return getIPError(myIP.IP, err)
}

// 从189云获取IP地址
func getMyIP189() (string, error) {
	// 获取IP地址
	var myIP struct {
		Data struct {
			IP string `json:"ip"`
		} `json:"data"`
	}
	_, err := resty.New().R().SetResult(&myIP).Get("https://b.cloud.189.cn/getWebImUrl.action")
	return getIPError(myIP.Data.IP, err)
}

func getMyIPiQiyi() (string, error) {
	res, err := resty.New().R().Get("https://data.video.iqiyi.com/v.f4v")
	if err != nil {
		return "", err
	}
	str := res.String()
	ip := IPRegexRule.FindString(str)
	return getIPError(ip, nil)
}

func getMyIPQQ() (string, error) {
	var myIP struct {
		IP string `json:"ip"`
	}

	_, err := resty.New().R().SetResult(&myIP).Get("https://r.inews.qq.com/api/ip2city?otype=json")
	return getIPError(myIP.IP, err)
}
