package test

import (
	"testing"
	"../component"
)

func TestWalkTorrent(t *testing.T) {
	path := "E:\\DP\\Sub"
	url := "http://dp:820425@wwww.sina.com.cn:6800/jsonrpc"
	component.WalkTorrent(path, url)
}


func TestRequestAuth(t *testing.T) {
	url := "http://dp:820425@192.168.2.22.2:6800/jsonrpc"
	component.RequestAuth(url)
}

func TestMakeRequestBody(t *testing.T)  {
	content := "2131313"
	component.MakeRequestBody(content)
	
}