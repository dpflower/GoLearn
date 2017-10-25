package main

import (
	"testing"
)

func TestIsExistsSub(t *testing.T) {
	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	lang := "chn"
	isexists := isExistsSub(fileName, lang)
	logger.Info(isexists)
	t.Logf("")

}
