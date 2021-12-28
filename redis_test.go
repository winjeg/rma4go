package main

import "testing"

func TestPrintStat(t *testing.T) {
	printKeyStat("127.0.0.1", "", 6379)
}
