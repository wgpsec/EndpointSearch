package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wgpsec/EndpointSearch/define"
	"github.com/wgpsec/EndpointSearch/utils/Error"
)

// Config struct is a wrapper of viper

// GlobalConfig default Global Variable for Config
var GlobalConfig *viper.Viper
var C define.Configure
var R define.Rules

// SpecificInit func is Init with specific Config file
func InitConfigure(file string) {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile(file)
	GlobalConfig.SetConfigType("yaml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		SaveConfig(file)
		Error.HandleFatal(fmt.Errorf("不存在 config.yaml，已生成"))
	}
	Error.HandleError(GlobalConfig.Unmarshal(&C))
}

// SpecificInit func is Init with specific Rule file
func InitRule(file string) {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile(file)
	GlobalConfig.SetConfigType("yaml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		SaveRule(file)
		Error.HandleFatal(fmt.Errorf("不存在 rule.yaml，已生成"))
	}
	Error.HandleError(GlobalConfig.Unmarshal(&R))
}

func SaveConfig(file string) {
	GlobalConfig.Set("CloudService", "oss,ecs")
	GlobalConfig.Set("Mode", ".")
	GlobalConfig.Set("Mode2", "-,.")
	GlobalConfig.Set("PortList", "80,443")
	GlobalConfig.Set("Prefix", "sonic,legacy,preprod,gamma,beta,staging")
	GlobalConfig.Set("Suffix", "sonic,legacy,preprod,gamma,beta,staging")
	GlobalConfig.SetConfigFile(file)
	err := GlobalConfig.WriteConfig()
	Error.HandleFatal(err)
}

func SaveRule(file string) {
	defaultRule :=
		[]define.RuleText{
			{
				Header: []string{"text/xml", "application/xml"},
				Body:   []string{"InvalidVersion"},
			},
		}
	GlobalConfig.Set("rules", defaultRule)
	GlobalConfig.SetConfigFile(file)
	err := GlobalConfig.WriteConfig()
	Error.HandleFatal(err)
}
