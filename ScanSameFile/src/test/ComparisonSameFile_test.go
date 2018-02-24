package test

import ("testing"
"../component"

)
func TestComparison(t *testing.T) {
	fullPath :="E:\\DP\\sub"
	component.Comparison(fullPath)
}


func TestConcurrentComparison(t *testing.T) {
	fullPath :="E:\\DP\\sub"
	component.ConcurrentComparison(fullPath)
}