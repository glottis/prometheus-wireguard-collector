package main

import (
	"testing"
)

func Test_parseDump(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s  string
		hs float64
		rx float64
		tx float64
	}{
		{"valid", "122q2qAC0ExW/BuGSmXT6g9Wd3oEGy5lA=\t1120u8LWR682knVm262z9Nf96P+m8=\t51820\toff\n923U/iBdcz8BcqE3Yo1pEHxBe+pdidQyB1=\t(none)\t1.1.3.1:51820\t10.90.0.12/32,10.0.3.0/24\t443\t12345\t1337\toff\n", 443, 12345, 1337},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDump(tt.s)
			if err != nil {
				t.Fatalf("parseDump failed, due to: %v\n", err)
			}
			for _, k := range got {
				if k.LatestHS != tt.hs {
					t.Fail()
					t.Logf("wanted %v as latest handshake but got %v\n", tt.hs, k.LatestHS)
				}
				if k.Rx != tt.rx {
					t.Fail()
					t.Logf("wanted %v as rx but got %v\n", tt.rx, k.Rx)
				}
				if k.Tx != tt.tx {
					t.Fail()
					t.Logf("wanted %v as tx but got %v\n", tt.tx, k.Tx)
				}
			}
		})
	}
}
