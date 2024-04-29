package pkg

import (
	"fmt"
	"github.com/wgpsec/EndpointSearch/define"
	"net"
	"net/http"
	"strings"
	"sync"
)

type ResponseData struct {
	url  string
	resp *http.Response
}

func SearchDomain(domainList ...string) (ipRecordList []define.IPRecord) {
	if len(domainList) != 0 {
		var answers []net.IP
		for _, domain := range domainList {
			answers, _ = net.LookupIP(domain)
			if len(answers) != 0 {
				ipRecordList = append(ipRecordList, define.IPRecord{Domain: domain, Ip: answers})
			}
		}
	}
	return ipRecordList
}

func SearchSRVRecord(ipRecordList ...define.IPRecord) (recordList []define.Record) {
	if len(ipRecordList) != 0 {
		for _, ipRecord := range ipRecordList {
			_, srv, err := net.LookupSRV(ipRecord.Domain, "tcp", "")
			if err != nil {
				recordList = append(recordList, define.Record{Ip: ipRecord.Ip, SvcDomain: ipRecord.Domain})
			} else {
				recordList = append(recordList, define.Record{Ip: ipRecord.Ip, SvcDomain: ipRecord.Domain, SrvRecords: srv})
			}
		}
	}
	return recordList
}

func SearchEndpoint(client *http.Client, portList []string, recordList ...define.Record) (resultList []string) {
	if len(recordList) != 0 {
		resultsChan := make(chan ResponseData, cap(recordList))
		var wg sync.WaitGroup
		for _, record := range recordList {
			for _, port := range portList {
				if len(record.SrvRecords) != 0 {
					for _, srv := range record.SrvRecords {
						wg.Add(1)
						go func(srv *net.SRV, port string, wg *sync.WaitGroup) {
							defer wg.Done()
							requestStr := strings.Join([]string{srv.Target, ":", fmt.Sprintf("%v", srv.Port)}, "")
							endpoint, resp := SendHttpRequest(client, requestStr)
							if endpoint != "" {
								resultsChan <- ResponseData{endpoint, resp}
							}

						}(srv, port, &wg)

					}
				}
				wg.Add(1)
				go func(record define.Record, port string, wg *sync.WaitGroup) {
					defer wg.Done()
					requestStr := strings.Join([]string{record.SvcDomain, ":", port}, "")
					endpoint, resp := SendHttpRequest(client, requestStr)
					if endpoint != "" {
						resultsChan <- ResponseData{endpoint, resp}
					}
				}(record, port, &wg)
			}
		}
		wg.Wait()
		close(resultsChan)

		for data := range resultsChan { // 从channel中收集结果
			if JudgeEndpoint(data.resp) {
				resultList = append(resultList, data.url)
			}
		}
	}
	return resultList
}

func SendHttpRequest(client *http.Client, request string) (endpoint string, resp *http.Response) {
	requestStr := strings.Join([]string{"http://", request}, "")
	resp, err := client.Get(requestStr)
	if err != nil {
		requestStr = strings.Join([]string{"https://", request}, "")
		resp, err = client.Get(requestStr)
		if err != nil {
			return "", resp
		}
	}
	return requestStr, resp
}
