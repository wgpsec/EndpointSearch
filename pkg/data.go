package pkg

import (
	"bufio"
	"github.com/wgpsec/EndpointSearch/internal/config"
	"os"
	"regexp"
	"strings"

	"github.com/wgpsec/EndpointSearch/utils/Error"
)

func ParseFileParameter(fileName string) (fileHostList []string) {
	file, err := os.Open(fileName)
	Error.HandlePanic(err)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		fileHostList = append(fileHostList, line)
	}
	file.Close()
	return fileHostList
}

func ConvertToReqList(service string, prefix string, suffix string, param ...string) (reqList []string) {
	if len(param) != 0 {
		for _, host := range param {
			// 当输入Url时提取出域名
			re := regexp.MustCompile(`(http|https)://`)
			if re.MatchString(host) {
				host = re.ReplaceAllString(host, "")
				re = regexp.MustCompile(`([/\\]).*`)
				if re.MatchString(host) {
					host = re.ReplaceAllString(host, "")
				}
			}
			// 匹配IP/域名
			if regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-_]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,11}$`).MatchString(host) {
				reqList = append(reqList, GenerateReqString(service, prefix, suffix, host)...)
			}
		}
	}
	return reqList
}

func GenerateReqString(service string, prefix string, suffix string, host string) (reqList []string) {
	modeList := strings.Split(config.C.Mode, ",")
	mode2List := strings.Split(config.C.Mode2, ",")
	serviceList := strings.Split(service, ",")
	prefixList := strings.Split(prefix, ",")
	suffixList := strings.Split(suffix, ",")
	var reqString string
	for _, mode := range modeList {
		for _, serviceName := range serviceList {
			for _, prefixName := range prefixList {
				if len(prefixList) == 0 {
					break
				}
				for _, mode2 := range mode2List {
					reqString = strings.Join([]string{prefixName, mode2, serviceName, mode, host}, "")
					reqList = append(reqList, reqString)
				}
			}
			for _, suffixName := range suffixList {
				if len(suffixList) == 0 {
					break
				}
				for _, mode2 := range mode2List {
					reqString = strings.Join([]string{serviceName, mode2, suffixName, mode, host}, "")
					reqList = append(reqList, reqString)
				}
			}
			reqString = strings.Join([]string{serviceName, mode, host}, "")
			reqList = append(reqList, reqString)
		}
	}
	return reqList
}
