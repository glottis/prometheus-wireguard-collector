package main

import (
	"log"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
)

type WgCollector struct {
	hs     *prometheus.Desc
	rx     *prometheus.Desc
	tx     *prometheus.Desc
	err    *prometheus.Desc
	device string // wg interface
}

// Returns a new collector with descriptors for named device and instance
func NewWgCollector(device string) *WgCollector {

	labels := prometheus.Labels{"device": device}

	rx := prometheus.NewDesc("wireguard_peer_transfer_rx_bytes_total", "Bytes received from Wireguard peer", []string{"peer"}, labels)
	tx := prometheus.NewDesc("wireguard_peer_transfer_tx_bytes_total", "Bytes sent to Wireguard peer", []string{"peer"}, labels)
	hs := prometheus.NewDesc("wireguard_peer_latest_handshake", "Latest handshake for Wireguard peer", []string{"peer"}, labels)
	err := prometheus.NewDesc("wireguard_peer_metric_collector_error", "Error metric for the Wireguard peer collector", nil, labels)

	return &WgCollector{
		device: device,
		rx:     rx,
		tx:     tx,
		hs:     hs,
		err:    err,
	}

}

// Describe sends metric descriptors over the channel
func (wgc *WgCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- wgc.hs
	ch <- wgc.rx
	ch <- wgc.tx
	ch <- wgc.err

}

// Collect fetches values and sends metrics over the channel
func (wgc *WgCollector) Collect(ch chan<- prometheus.Metric) {

	cmd := exec.Command("wg", "show", wgc.device, "dump")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running wg show dump cmd: %v\n", err)
		ch <- prometheus.MustNewConstMetric(wgc.err, prometheus.CounterValue, float64(1))
		return
	}

	peers, err := parseDump(string(stdoutStderr))
	if err != nil {
		log.Printf("Error parsing wg show dump cmd: %v\n", err)
		ch <- prometheus.MustNewConstMetric(wgc.err, prometheus.CounterValue, float64(1))
		return
	}

	for peer, v := range peers {
		ch <- prometheus.MustNewConstMetric(wgc.hs, prometheus.CounterValue, v.LatestHS, peer)
		ch <- prometheus.MustNewConstMetric(wgc.rx, prometheus.CounterValue, v.Rx, peer)
		ch <- prometheus.MustNewConstMetric(wgc.tx, prometheus.CounterValue, v.Tx, peer)
	}

	ch <- prometheus.MustNewConstMetric(wgc.err, prometheus.CounterValue, float64(0))
}
