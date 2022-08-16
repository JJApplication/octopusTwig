/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"errors"

	"github.com/JJApplication/fushin/server/uds"
)

// 转发接口请求

func Proxy() uds.Func {
	return func(c *uds.UDSContext, req uds.Req) {
		// 抛弃无操作
		if req.Operation == "" {
			err := c.Response(uds.Res{
				Error: errors.New(ErrNoOp).Error(),
				Data:  "",
				From:  OctopusTwig,
				To:    nil,
			})
			if err != nil {
				logger.ErrorF("proxy [empty operation] error: %s", err.Error())
			}
			return
		}

		// 抛弃无属主的请求
		if req.From == "" {
			err := c.Response(uds.Res{
				Error: errors.New(ErrNoFrom).Error(),
				Data:  "",
				From:  OctopusTwig,
				To:    nil,
			})
			if err != nil {
				logger.ErrorF("proxy [unknown address] error: %s", err.Error())
			}
			return
		}

		// 无目标服务时
		if req.To == nil {
			err := c.Response(uds.Res{
				Error: errors.New(ErrNoTo).Error(),
				Data:  "",
				From:  OctopusTwig,
				To:    nil,
			})
			if err != nil {
				logger.ErrorF("proxy [empty address] error: %s", err.Error())
			}
			return
		}

		// 开始转发 目标为自身的请求忽略
		sender(c, req)
	}
}

// 转发器 转发原样的报文到指定服务
// 默认写入返回值
// 因为fushin的限制 不支持的operation在proxy端仅记录报错日志 不会返回错误
func sender(c *uds.UDSContext, req uds.Req) {
	for _, t := range req.To {
		toAddress := GetSocket(t)
		if toAddress == "" {
			continue
		}
		logger.InfoF("proxy [sender] -> %s", t)
		cli := GetClientFromPool(t)
		_ = cli.Dial()
		res, err := cli.SendWithRes(uds.Req{
			Operation: req.Operation,
			Data:      req.Data,
			From:      req.From,
			To:        req.To,
		})
		if err != nil {
			logger.ErrorF("proxy [sender] -> %s error: %s", t, err.Error())
		}
		logger.InfoF("proxy [sender] -> %s done %+v", t, res)
		c.Response(res)
		PutClientToPool(cli)
	}
}
