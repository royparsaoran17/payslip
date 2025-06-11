// Package bootstrap
package bootstrap

import (
	"payroll-se/internal/consts"
	"payroll-se/pkg/logger"
	"payroll-se/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
