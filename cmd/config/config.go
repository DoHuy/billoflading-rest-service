package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"

	commonConf "vtp-apis/packages/config"
)

var (
	c    Conf
	once sync.Once
)

type Conf struct {
	RunMode 	commonConf.RunModeConfig
	HTTP    	commonConf.HTTPConfig
	Elastic	    commonConf.ElasticConfig
}
func Config() *Conf {
	once.Do(initConfig)
	return &c
}
func initConfig() {
	c = Conf{}
	envconfig.MustProcess("", &c.RunMode)
	envconfig.MustProcess("", &c.HTTP)
	envconfig.MustProcess("", &c.Elastic)
}
