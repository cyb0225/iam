/**
@author: yeebing
@date: 2022/9/27
**/

package apiserver

import (
	"fmt"
	"github.com/cyb0225/iam/pkg/db"
	"github.com/cyb0225/iam/pkg/email"
	"github.com/cyb0225/iam/pkg/log"
	"github.com/cyb0225/iam/pkg/server"
	"github.com/spf13/viper"
)

// Option stored the whole options that apiserver needs.
type Option struct {
	Mysql db.Option `yaml:"mysql"`
	//Redis  redis.Option  `yaml:"redis"`
	Log    log.Option    `yaml:"log"`
	Email  email.Option  `yaml:"email"`
	Server server.Option `yaml:"server"`
}

var (
	Opts *Option
)

// NewOption read config from config file.
// input the config file's path.
func NewOption(config string) (*Option, error) {
	opts := &Option{}

	vp := viper.New()
	if len(config) != 0 {
		vp.AddConfigPath(config)
	}
	vp.AddConfigPath(".")
	vp.AddConfigPath("./config")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	vp.WatchConfig()
	vp.OnConfigChange(nil)

	if err := vp.Unmarshal(opts); err != nil {
		return nil, err
	}

	Opts = opts

	fmt.Printf("%#+v", *opts)
	return opts, nil
}
