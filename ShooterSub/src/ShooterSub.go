package src

import (
	"./subcomponent"
	"fmt"
	logging "github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("submon")
)

func main() {
	fmt.Println("Hello World!")

	fileName := "E:\\DP\\CentOS-7.0-1406-x86_64-Everything.iso"
	fmt.Println(subcomponent.Getfilehash(fileName))

}
