package services

import "session15/storage"

func SearchServices(stringdata []string, baseurl string) {
	for _, str := range stringdata {
		//	storage.Mappp[str][baseurl] = storage.Mappp[str][baseurl] + 1
		//}
		if storage.Mappp[str] == nil {
			storage.Mappp[str] = make(map[string]int)
		}
		storage.Mu.Lock()
		storage.Mappp[str][baseurl]++
		storage.Mu.Unlock()
	}
}
