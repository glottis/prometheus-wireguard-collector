package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	i string // interface
	l string // listen
	p string // port
)

func parseFlags() {
	flag.StringVar(&i, "i", "wg0", "wireguard interface")
	flag.StringVar(&l, "l", "*", "address to listen on for polling")
	flag.StringVar(&p, "p", "8090", "listening port")
	flag.Parse()

}

func main() {

	parseFlags()

	collector := NewWgCollector(i)
	prometheus.MustRegister(collector)

	// Expose metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf("%s:%s", l, p)
	log.Printf("listening on %s\n", addr)

	http.ListenAndServe(addr, nil)

}
