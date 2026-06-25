## Simple Prometheus Wireguard exporter

Exports the following metrics for each wg peer for a wg interface:
  - latest handshake
  - transfer rx
  - transfer tx
  - error (if something goes wrong collecting metrics)

As it is stateless there is no need to restart the collector if adding new peers

### Buildning/testing:
  - go test ./...
  - go build -o wg-exporter main.go

### Flags:
  - i : wireguard interface
  - l : listen address
  - p : port to use

License: MIT
