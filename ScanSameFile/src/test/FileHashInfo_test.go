package test

import (
	"testing"
	"../component"
	"github.com/labstack/gommon/log"
	"time"
)

func TestGetFileHashDetailInfo(t *testing.T) {
	fullFileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	begin := time.Now()
	info, err := component.GetFileHashDetailInfo(fullFileName)

	if err != nil {
		log.Error(err)
	}
	duration := time.Now().Sub(begin)
	log.Info(duration)
	log.Info(info)
}
