package services

import (
	"csvAggregator/internals/models"
	"log"
)

func SplitFileToN(Filedata [][]string) {

	lenngth := len(Filedata) / models.N
	for i := 0; i < models.N; i++ {
		if i == models.N-1 {
			models.SplitFileMapp[i] = Filedata[i*lenngth : len(Filedata)]
		} else {
			models.SplitFileMapp[i] = Filedata[i*lenngth : (i+1)*lenngth]
		}
	}
	log.Println(models.SplitSuccessFull)

}
