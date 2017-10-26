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

func TestWalkDir(t *testing.T) {
	fileName := "E:\\DP"
	lang := "chn"
	walkDir(fileName, lang)

	t.Logf("")

}
