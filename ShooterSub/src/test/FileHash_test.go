package test

import (
	"../subcomponent"
	"fmt"
	"testing"
)

func TestComputeFileHash(t *testing.T) {
	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	fmt.Println(subcomponent.Getfilehash(fileName))

	//t.Logf("11")
}
