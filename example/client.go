package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/devsisters/goquic"
	"github.com/devsisters/gospdyquic"
)

//var host string
var url string
var logLevel int

func init() {
	//	flag.StringVar(&host, "host", "127.0.0.1:8080", "host to connect")
	flag.StringVar(&url, "url", "http://127.0.0.1:8080/", "host to connect")
	flag.IntVar(&logLevel, "loglevel", -1, "Log level")
}

func main() {
	goquic.Initialize()
	goquic.SetLogLevel(logLevel)

	flag.Parse()

	client := &http.Client{
		Transport: &gospdyquic.QuicRoundTripper{},
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
