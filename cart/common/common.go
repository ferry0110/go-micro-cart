package common

import (
	"encoding/json"
)

// SwapTo 通过json tag 进行数据的转移
func SwapTo(sourceData, targetData interface{}) (err error)   {
	dataByte,err := json.Marshal(sourceData)
	if err!=nil {
		return
	}
	return json.Unmarshal(dataByte,targetData)
}
