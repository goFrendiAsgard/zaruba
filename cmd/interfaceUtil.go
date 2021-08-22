package cmd

import (
	"encoding/json"
	"fmt"
)

func printInterface(obj interface{}) {
	if obj == nil {
		fmt.Println("")
		return
	}
	jsonByte, err := json.Marshal(obj)
	if err == nil && len(jsonByte) > 0 {
		ch := jsonByte[0]
		if ch == []byte("{")[0] || ch == []byte("[")[0] {
			fmt.Println(string(jsonByte))
			return
		}
	}
	fmt.Println(obj)
}
