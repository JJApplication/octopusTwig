/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"fmt"
	"path"

	"github.com/JJApplication/fushin/utils/env"
)

const (
	OctopusTwig = "Twig"
	// SocketRoot env
	SocketRoot = "SocketRoot"
	// UnixSocketAddress 监听地址
	UnixSocketAddress = "UnixSocketAddress"
)

var envLoader = env.EnvLoader{}

func GetSocket(app string) string {
	if app == "" {
		return ""
	}
	return path.Join(envLoader.Get(SocketRoot).Raw(), fmt.Sprintf("%s.sock", app))
}

func GetListenAddr() string {
	return envLoader.Get(UnixSocketAddress).Raw()
}
