package gconfig

import (
	"go.uber.org/zap"

	"github.com/kelseyhightower/envconfig"
)

// GConfig represents Global Config applied to all applications. All settings should have default value.
type GConfig struct {
	DisableSampleZapLog bool `envconfig:"DISABLE_SAMPLE_ZAP_LOG" default:"true"`
}

var (
	gConf GConfig
)

func init() {
	if err := envconfig.Process("", &gConf); err != nil {
		zap.S().Panic("load global config error: ", err.Error())
	}
}

func GlobalConfig() GConfig {
	return gConf
}
