package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cocafe/readAbiFile/common"
)

func GetContractFunc(j *common.JsonBodyArr, f string) (common.JsonBody, bool) {
	var conFun common.JsonBody
	for i := 0; i < len(j.JsonArr); i++ {
		if j.JsonArr[i].Name == f {
			conFun = j.JsonArr[i]
			return conFun, true
		}
	}
	return conFun, false
}

func ParseData(abiPath string) *common.JsonBodyArr {
	var jsonBodyArr *common.JsonBodyArr
	f, err := os.OpenFile(abiPath, os.O_RDONLY, 0600)
	var contentByte []byte
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		contentByte, err = ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	tokenStr := "{\"token\":" + string(contentByte) + "}"
	err = json.Unmarshal([]byte(tokenStr), &jsonBodyArr)
	if err != nil {
		fmt.Println("unmarshal error")
		return nil
	}
	return jsonBodyArr
}
