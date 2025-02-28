package get

import (
	"fmt"
	"session-revision/put"
)

var size = put.Size

func Getting(key int) {
	fmt.Println("getting...")
	_, ok := put.Mp[key]
	if ok {
		fmt.Println(put.Mp[key])
		for i := 0; i < len(put.Trace); i++ {
			if put.Trace[i] == key {
				put.Trace = append(put.Trace[:i], put.Trace[i+1:]...)
				put.Trace = append(put.Trace, key)
				break
			}
		}
	} else {
		fmt.Println("-1")
	}

}
