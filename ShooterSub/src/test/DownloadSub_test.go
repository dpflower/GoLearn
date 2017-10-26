package test

import (
	"../subcomponent"
	"testing"
)

func TestDownloadSub(t *testing.T) {
	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	lang := "chn"
	subcomponent.DownloadSub(fileName, lang)
	t.Logf("")

}
