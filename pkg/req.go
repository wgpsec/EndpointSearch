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

func SearchDomain(domainList ...string) (ipRecordList []IPRecord) {
	if len(domainList) != 0 {
		var answers []net.IP
		for _, domain := range domainList {
			answers, _ = net.LookupIP(domain)
			if len(answers) != 0 {
				ipRecordList = append(ipRecordList, IPRecord{domain: domain, ip: answers})
			}
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
				recordList = append(recordList, Record{ip: ipRecord.ip, svcDomain: ipRecord.domain, srvRecords: srv})
			}
		}
	}
	return recordList
}

func SearchEndpoint(client *http.Client, portList []string, recordList ...Record) (respList []ResponseData) {
	if len(recordList) != 0 {
		resultsChan := make(chan ResponseData, cap(recordList))
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		for _, record := range recordList {
			for _, port := range portList {
				if len(record.srvRecords) != 0 {
					resultsChan2 := make(chan ResponseData, cap(record.srvRecords))
					for _, srv := range record.srvRecords {
						go func(srv *net.SRV, port string, wg *sync.WaitGroup) {
							defer wg.Done()
							//logrus.Debug(srv.Target, ":", port, " Working")
							requestStr := strings.Join([]string{srv.Target, ":", fmt.Sprintf("%v", srv.Port)}, "")
							SendHttpRequest(client, requestStr, resultsChan2, &wg2)
							//logrus.Debug(srv.Target, ":", port, " Done")
						}(srv, port, &wg2)
						go func() {
							respList = append(respList, <-resultsChan)
						}()
					}
					wg2.Wait()
					close(resultsChan2)
				}
				wg1.Add(1)
				go func(record Record, port string, wg *sync.WaitGroup) {
					//logrus.Debug(record.SvcDomain, ":", port, " Working")
					requestStr := strings.Join([]string{record.svcDomain, ":", port}, "")
					SendHttpRequest(client, requestStr, resultsChan, wg)
					//logrus.Debug(record.SvcDomain, ":", port, " Done")
				}(record, port, &wg1)
				go func() {
					respList = append(respList, <-resultsChan)
				}()
			}
		}
		wg1.Wait()
		close(resultsChan)
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
		resp.Header.Clone()
		resp.Body.Close()
		if contentType != "" {
			resultsChan <- ResponseData{requestStr, contentType}
		}
	}
}
