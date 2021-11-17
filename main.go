package main

import (
	"flag"
	"fmt"

	"github.com/cocafe/readAbiFile/tool"
)

func main() {
	abiPath := flag.String("path", "./contract/token.abi", "contract abi path")
	abiMethod := flag.String("method", "Approval", "method in abi")
	flag.Parse()
	fmt.Println(*abiPath, *abiMethod)
	jsonBodyArr := tool.ParseData(*abiPath)
	if jsonBodyArr != nil {
		funcStruct, flag := tool.GetContractFunc(jsonBodyArr, *abiMethod)
		if !flag {
			fmt.Println("no have this function")
			return
		}
		fmt.Println(funcStruct)
	}
}
