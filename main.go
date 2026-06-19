package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&i, "i", "wg0", "which wireguard interface")
	flag.StringVar(&l, "l", "", "which address to listen on")
	flag.StringVar(&p, "p", "8090", "which port to use")
	flag.Parse()

}

func main() {

	parseFlags()

	collector := NewWgCollector(i)
	prometheus.MustRegister(collector)

	// Expose metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf("%s:%s", l, p)
	http.ListenAndServe(addr, nil)

}
