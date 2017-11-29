package test

import ("testing"
	"../component"
)

func TestWalkTorrent(t *testing.T) {
	path := "E:\\DP\\Sub"
	component.WalkTorrent(path)
}
