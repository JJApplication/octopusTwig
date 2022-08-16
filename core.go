/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"github.com/JJApplication/fushin/server/uds"
)

type TwigCore struct {
	Name string
	us   *uds.UDSServer
}

func (t *TwigCore) Init() {
	udsServer := &uds.UDSServer{
		Name:   GetListenAddr(),
		Option: uds.DefaultOption,
		Logger: nil,
	}
	udsServer.Option.AutoCheck = false
	udsServer.Option.MaxSize = DefaultBufSize
	t.us = udsServer
}

func (t *TwigCore) Start() {
	if err := t.us.Proxy(Proxy()); err != nil {
		logger.ErrorF("%s uds server start error: %s", t.Name, err.Error())
	}
}
