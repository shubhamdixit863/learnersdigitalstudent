package services

import (
	"csvAggregator/internals/models"
	"log"
	"strconv"
	"sync"
)

func ProcessData(ch chan map[string]float64, filedata [][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	mp := make(map[string]float64)
	for _, v := range filedata {
		fl, _ := strconv.ParseFloat(v[1], 64)
		mp[v[0]] = fl + mp[v[0]]
	}
	log.Println(models.Fileprocessed, "   ", mp)
	ch <- mp
}
