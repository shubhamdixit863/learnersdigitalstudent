package put

import (
	"fmt"
)

var Size = 5
var Trace []int
var Mp = make(map[int]string)

func Putting(key int, val string) {
	fmt.Println("putting...")
	_, ok := Mp[key]
	if ok {
		Mp[key] = val
		for i := 0; i < len(Trace); i++ {
			if Trace[i] == key {
				Trace = append(Trace[:i], Trace[i+1:]...)
				Trace = append(Trace, key)
				fmt.Println("trace is___ ", Trace)
				break
			}
		}
	} else {
		Mp[key] = val
		if len(Trace) > Size {
			Trace = Trace[1:]
		}
		Trace = append(Trace, key)
		fmt.Println("trace is ", Trace)
	}

}
