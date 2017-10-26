package test

import (
	"../subcomponent"
	"testing"
)

func TestWalkDir(t *testing.T) {
	fileName := "E:\\DP"
	lang := "chn"
	subcomponent.WalkDir(fileName, lang)

	t.Logf("")

}
