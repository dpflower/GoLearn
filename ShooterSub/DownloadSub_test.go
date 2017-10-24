package main

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	lang := "chn"
	downloadSub(fileName, lang)
	t.Logf("")

}
