package pkg

import (
	"net"
)

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
