package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Printf("\nExample-Command :  ./program_name  list_of_urls.csv  active_urls.txt\n\n")
		return
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Timeout: 5 * time.Second}
	httpsClient := &http.Client{Timeout: 5 * time.Second, Transport: tr}

	c := make(chan urlData)
	defer close(c)

	csvFile, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error Opening 'urls.csv':", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Read() // Skip Header

	nURLs := 0
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error Reading Record:", err)
			return
		}

		nURLs++
		go checkURL(httpsClient, httpClient, record, c)
	}

	f, _ := os.OpenFile(args[2], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer f.Close()

	var result urlData
	up := 0
	down := 0

	fmt.Printf("\nChecking %d URLs...\n\n", nURLs)

	for i := range nURLs {
		result = <-c
		if result.isValid {
			up++
			log.Printf("-- [%02d] --- UP --- %s\n", i+1, result.url)

			urlNumber := fmt.Sprintf("[%02d]", up)
			f.WriteString(urlNumber + " Category : " + result.category + "\n")
			f.WriteString("     URL      : " + result.url + "\n")
			f.WriteString("     Comment  : " + result.comment + "\n\n\n")
		} else {
			down++
			log.Printf("-- [%02d] -- DOWN -- %s\n", i+1, result.url)
		}
	}

	fmt.Printf("\nDone...\n\nServers UP   : %d\nServers DOWN : %d\n\nCheck File   : '%s'\n\n", up, down, args[2])
}
