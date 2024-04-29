package pkg

import (
	"crypto/tls"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"net/http"
	"os"
	"time"
)

func GetPwd() (homePath string) {
	homePath, err := os.Getwd()
	Error.HandlePanic(err)
	return homePath
}

func GenerateHTTPClient(timeOut int) *http.Client {
	client := &http.Client{
		Timeout: time.Duration(timeOut) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return client
}
