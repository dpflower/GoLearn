package test

import (
	"testing"
	"../component"
)

func TestWalkTorrent(t *testing.T) {
	path := "E:\\DP\\Sub"
	url := "http://dp:820425@192.168.2.222:6800/jsonrpc"
	component.WalkTorrent(path, url)
}


func TestRequestAuth(t *testing.T) {
	url := "http://dp:820425@192.168.2.22.2:6800/jsonrpc"
	component.RequestAuth(url)
}