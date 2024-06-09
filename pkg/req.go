package pkg

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
)

type ResponseData struct {
	url         string
	contentType string
}

type IPRecord struct {
	domain string
	ip     []net.IP
}

type Record struct {
	ip         []net.IP
	svcDomain  string
	srvRecords []*net.SRV
}

func SearchDomain(reqList ...string) (ipRecordList []IPRecord) {
	if len(reqList) != 0 {
		resultsChan := make(chan IPRecord, cap(reqList)*2)
		var wg sync.WaitGroup
		var answers []net.IP
		for _, domain := range reqList {
			wg.Add(1)
			go func(domain string, wg *sync.WaitGroup) {
				defer wg.Done()
				answers, _ = net.LookupIP(domain)
				if len(answers) != 0 {
					fmt.Printf("[+] %s\n", domain)
					resultsChan <- IPRecord{domain, answers}
				}
			}(domain, &wg)
		}
		wg.Wait()
		close(resultsChan)
		for data := range resultsChan {
			ipRecordList = append(ipRecordList, data)
		}
	}
	return ipRecordList
}

func SearchSRVRecord(ipRecordList ...IPRecord) (recordList []Record) {
	if len(ipRecordList) != 0 {
		for _, ipRecord := range ipRecordList {
			_, srv, err := net.LookupSRV(ipRecord.domain, "tcp", "")
			if err != nil {
				recordList = append(recordList, Record{ip: ipRecord.ip, svcDomain: ipRecord.domain})
			} else {
				for _, srvRecord := range srv {
					fmt.Printf("[+] %s:%v\n", srvRecord.Target, srvRecord.Port)
				}
				recordList = append(recordList, Record{ip: ipRecord.ip, svcDomain: ipRecord.domain, srvRecords: srv})
			}
		}
	}
	return recordList
}

func SearchEndpoint(client *http.Client, portList []string, recordList ...Record) (respList []ResponseData) {
	if len(recordList) != 0 {
		resultsChan := make(chan ResponseData, cap(recordList)*cap(portList)*10)
		var wg sync.WaitGroup
		for _, record := range recordList {
			for _, port := range portList {
				if len(record.srvRecords) != 0 {
					for _, srv := range record.srvRecords {
						wg.Add(1)
						go func(srv *net.SRV, wg *sync.WaitGroup) {
							requestStr := strings.Join([]string{srv.Target, ":", fmt.Sprintf("%v", srv.Port)}, "")
							SendHttpRequest(client, requestStr, resultsChan, wg)
						}(srv, &wg)
					}
				}
				wg.Add(1)
				go func(record Record, port string, wg *sync.WaitGroup) {
					requestStr := strings.Join([]string{record.svcDomain, ":", port}, "")
					SendHttpRequest(client, requestStr, resultsChan, wg)
				}(record, port, &wg)
			}
		}
		wg.Wait()
		close(resultsChan)
		for resp := range resultsChan {
			respList = append(respList, resp)
		}
	}
	return respList
}

func SendHttpRequest(client *http.Client, request string, resultsChan chan ResponseData, wg *sync.WaitGroup) {
	defer wg.Done()
	requestStr := strings.Join([]string{"http://", request}, "")
	resp, err := client.Get(requestStr)
	if err != nil {
		requestStr = strings.Join([]string{"https://", request}, "")
		resp, err = client.Get(requestStr)
		if err != nil {
			return
		}
	}
	if resp != nil {
		contentType := resp.Header.Get("Content-Type")
		resp.Body.Close()
		if contentType != "" {
			resultsChan <- ResponseData{requestStr, contentType}
		}
	}
}
