package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Wireguard peer (not all) fields
type Peer struct {
	LatestHS float64 // latest handshake
	Rx       float64 // transfer rx in bytes
	Tx       float64 // transfer tx in bytes
}

// returns the output from wg dump in a parsed way
func parseDump(dumpOutput string) (map[string]Peer, error) {
	m := map[string]Peer{}
	dumpOutput = strings.TrimSpace(dumpOutput)
	peers := strings.Split(dumpOutput, "\n")

	if len(peers) <= 1 {
		return m, nil
	}

	for _, row := range peers[1:] {
		v := strings.Split(row, "\t")
		length := len(v)
		if length < 8 {
			return m, fmt.Errorf("invalid size of array: %v", length)
		}

		handShake, err := strconv.ParseFloat(v[4], 64)
		if err != nil {
			return m, fmt.Errorf("failed to parse latest handshake field: %v", err)
		}
		rx, err := strconv.ParseFloat(v[5], 64)
		if err != nil {
			return m, fmt.Errorf("failed to parse rx field: %v", err)
		}
		tx, err := strconv.ParseFloat(v[6], 64)
		if err != nil {
			return m, fmt.Errorf("failed to parse tx field: %v", err)
		}

		m[v[0]] = Peer{LatestHS: handShake, Rx: rx, Tx: tx}
	}

	return m, nil
}
