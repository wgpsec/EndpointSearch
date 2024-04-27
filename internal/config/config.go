package config

import (
	"github.com/spf13/viper"
	"github.com/wgpsec/EndpointSearch/define"
	"github.com/wgpsec/EndpointSearch/utils/Error"
)

// Config struct is a wrapper of viper

// GlobalConfig default Global Variable for Config
var GlobalConfig *viper.Viper
var C define.Configure

// SpecificInit func is Init with specific Config file
func Init(file string) {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile(file)
	GlobalConfig.SetConfigType("json")
	err := GlobalConfig.ReadInConfig()
	Error.HandleFatal(err, "Config File "+file+" can't read or not exist.")
	Error.HandleError(GlobalConfig.Unmarshal(&C))
}

func SaveConfig() {
	// GlobalConfig.Set("log_level", C.LogLevel)
	err := GlobalConfig.WriteConfig()
	Error.HandleFatal(err)
}
