package test

import (
	"../subcomponent"
	"testing"
)

func TestSubHDDownload(t *testing.T) {
	fileName := "E:\\DP\\sub\\The.Orville.S01E05.720p.HDTV.x264-AVS.mkv"
	lang := "chn"
	subcomponent.SubHDDownload(fileName, lang)

	//t.Logf("11")
}