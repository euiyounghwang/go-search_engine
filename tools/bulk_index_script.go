package main

import (
	"flag"
	"log"
)

var (
	es_host  string
)

func init() {
	// flag.IntVar(es_host, "es_host", "http://localhost:9209", "Host target")
	flag.StringVar(&es_host, "es_host", "http://localhost:9209", "Host target")
	flag.Parse()
}

func main() {
	// log.Println("main")
	log.Println(es_host)
}