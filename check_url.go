package main

import (
	"log"
	"net/http"
	"net/url"
)

type urlData struct {
	category string
	url      string
	comment  string
	isValid  bool
}

func checkURL(httpsClient *http.Client, httpClient *http.Client, record []string, c chan urlData) {
	var urlHTTP, urlHTTPS, URL string

	u, err := url.ParseRequestURI(record[1])
	if err != nil || u.Scheme == "" || u.Host == "" {
		log.Println("Invalid URL: ", record[1])
		c <- urlData{
			category: record[0],
			url:      record[1],
			comment:  record[2],
			isValid:  false,
		}
		return
	} else {
		urlHTTP = "http://" + u.Host + u.Path
		urlHTTPS = "https://" + u.Host + u.Path
	}

	_, errHTTPS := httpsClient.Get(urlHTTPS)
	_, errHTTP := httpClient.Get(urlHTTP)

	if errHTTPS == nil {
		URL = urlHTTPS
	} else {
		URL = urlHTTP
	}

	c <- urlData{
		category: record[0],
		url:      URL,
		comment:  record[2],
		isValid:  errHTTP == nil || errHTTPS == nil,
	}
}
