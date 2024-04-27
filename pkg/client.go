package pkg

import (
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"os"
)

func GetPwd() (homePath string) {
	homePath, err := os.Getwd()
	Error.HandlePanic(err)
	return homePath
}
