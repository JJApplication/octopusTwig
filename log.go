/*
Create: 2022/8/17
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopusTwig
package octopusTwig

import (
	"github.com/JJApplication/fushin/log"
)

var logger log.Logger

func init() {
	logger = log.Logger{
		Name:   OctopusTwig,
		Option: log.DefaultOption,
		Sync:   true,
	}
	logger.Init()
}
