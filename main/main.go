package main

import (
	"errors"
	"fmt"
	"github.com/wgpsec/EndpointSearch/cmd"
	"github.com/wgpsec/EndpointSearch/internal/config"
	"github.com/wgpsec/EndpointSearch/pkg"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"github.com/wgpsec/EndpointSearch/utils/File"
	"strings"
)

func init() {
	configFile := pkg.GetPwd()
	configFile = strings.Join([]string{configFile, "/config.json"}, "")
	_, err := File.FileCreateIfNonExist(configFile)
	Error.HandleFatal(err)
	config.Init(configFile)
	if config.C.CloudEndpoint == "" || config.C.Mode == "" {
		Error.HandleFatal(errors.New("请配置config.json"))
		return
	}
}

func main() {
	fmt.Println(cmd.RootCmd.Long)
	cmd.Execute()
}
