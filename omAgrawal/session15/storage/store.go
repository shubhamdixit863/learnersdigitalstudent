package storage

import "sync"

var Mappp = make(map[string]map[string]int)
var BaseUrl string
var VisitedUrl = make(map[string]int)

var UrlSlice = make([]string, 0)

var UrlChan = make(chan string)
var Mu sync.Mutex
var Wg sync.WaitGroup

//UrlSlice[0]="https://usf-cs272-s25.github.io/top10/"
//
//UrlSlice
