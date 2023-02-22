package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMx, hasSPF, spfRecord, HasDMARC, dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println("error in reading : ", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDmarc bool
	var spfRecord, dmarcRecord string
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Println("error in reading the domain ", err)

	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("error in txtRecord ", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Println("err in dmarc record ", err)
	}

	for _, dRecord := range dmarcRecords {
		if strings.HasPrefix(dRecord, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = dRecord
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMx, hasSPF, spfRecord, hasDmarc, dmarcRecord)

}
