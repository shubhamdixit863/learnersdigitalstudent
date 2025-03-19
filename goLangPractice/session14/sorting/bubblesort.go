package sorting

import (
	"fmt"
	"sync"
)

func BubbleSort(slc []int) {
	for i := 0; i < len(slc); i++ {
		for j := 0; j < len(slc)-1; j++ {
			//Swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}
	}
}

func BubbleOuter(slc []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(slc); i++ {
		ch <- i
	}

	close(ch) // This tells the range that the outer loop stops sending data
}

func BubbleInner(slc []int, ch chan int, wg *sync.WaitGroup) {
	//_, ok := <-ch
	//if !ok {
	//	defer wg.Done()
	//}

	defer wg.Done()
	for v := range ch {
		fmt.Println(v)

		for j := 0; j < len(slc)-1; j++ {
			//Swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}
	}

}

func BubbleOuterBuf(slc []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(slc); i++ {
		ch <- i
	}

	close(ch) // This tells the range that the outer loop stops sending data
}

func BubbleInnerBuf(slc []int, ch chan int, wg *sync.WaitGroup) {
	var ok = true
	defer wg.Done()

	for ok {
		v, ok := <-ch
		fmt.Println(ok, v)
		for j := 0; j < len(slc)-1; j++ {
			//Swap
			if slc[j] > slc[j+1] {
				temp := slc[j+1]
				slc[j+1] = slc[j]
				slc[j] = temp
			}
		}
	}

}
