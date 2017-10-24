package main

import (
	"fmt"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	fmt.Println(computeFileHash(fileName))

	t.Logf(computeFileHash(fileName))
}
