/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"github.com/JJApplication/fushin/client/uds"
)

// 生成全新的client
// 需要指定to Address

const (
	DefaultBufSize = 10 << 20
)

func NewClient(rec int) *uds.UDSClient {
	if rec <= 0 {
		rec = DefaultBufSize
	}
	return &uds.UDSClient{
		Addr:        "",
		MaxRecvSize: rec,
	}
}

// 断言client对象
func assertClient(c interface{}) (*uds.UDSClient, bool) {
	if v, ok := c.(*uds.UDSClient); ok {
		return v, true
	}
	return nil, false
}
