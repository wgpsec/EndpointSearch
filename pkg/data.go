package pkg

import (
	"bufio"
	"fmt"
	"github.com/wgpsec/EndpointSearch/internal/config"
	"github.com/wgpsec/EndpointSearch/utils/Error"
	"os"
	"regexp"
	"strings"
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

func ConvertToReqList(endpoint string, param ...string) (reqList []string) {
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
				Modes := strings.Split(config.C.Mode, ",")
				Endpoints := strings.Split(endpoint, ",")

				for _, mode := range Modes {
					for _, endpoint := range Endpoints {
						reqString := strings.Join([]string{endpoint, mode, host}, "")
						reqList = append(reqList, reqString)
					}
				}
			}
		}
	}
	return reqList
}

func ParseRecordResult(recordList ...Record) (resultList []string) {
	if len(recordList) != 0 {
		for _, record := range recordList {
			if len(record.srvRecords) != 0 {
				for _, srv := range record.srvRecords {
					result := strings.Join([]string{srv.Target, ":", fmt.Sprintf("%v", srv.Port)}, "")
					resultList = append(resultList, result)
				}
			}
			resultList = append(resultList, record.svcDomain)
		}
	}
	return resultList
}
