/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"sync"

	"github.com/JJApplication/fushin/client/uds"
)

// 连接池 维护一个可复用的uds连接
// 因为uds的长连接特性 在定时器后主动关闭uds连接来刷新pool

const (
	N = 20
)

var TwigPool = sync.Pool{New: func() interface{} {
	return NewClient(0)
}}

// GetClientFromPool 寻找可以复用的client
func GetClientFromPool(App string) *uds.UDSClient {
	// 循环尝试拿取
	for range [N]struct{}{} {
		c := TwigPool.Get()
		uc, ok := assertClient(c)
		if !ok {
			continue
		}
		// 为空表示未初始化 直接初始化
		if uc.Addr == "" {
			uc.Addr = GetSocket(App)
			return uc
		}
		if uc.Addr == GetSocket(App) {
			return uc
		}
		continue
	}
	// 没有pool池数据 只能初始化
	uc := NewClient(DefaultBufSize)
	uc.Addr = GetSocket(App)
	return uc
}

// PutClientToPool 使用完毕后返回池内
func PutClientToPool(uc *uds.UDSClient) {
	TwigPool.Put(uc)
}
