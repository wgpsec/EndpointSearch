package main

import (
	"fmt"
	"github.com/wgpsec/EndpointSearch/cmd"
	"github.com/wgpsec/EndpointSearch/internal/config"
	"github.com/wgpsec/EndpointSearch/utils/Error"
)

func init() {
	config.InitConfigure("config.yaml")
	config.InitRule("rule.yaml")

}

func main() {
	if config.C.CloudService == "" || config.C.Mode == "" || config.C.PortList == "" || config.C.Mode2 == "" {
		Error.HandleFatal(fmt.Errorf("请配置 config.yaml"))
		return
	}
	if len(config.R.RuleText) == 0 {
		Error.HandleFatal(fmt.Errorf("请配置 rule.yaml"))
		return
	}

	fmt.Println(cmd.RootCmd.Long)
	cmd.Execute()
}
