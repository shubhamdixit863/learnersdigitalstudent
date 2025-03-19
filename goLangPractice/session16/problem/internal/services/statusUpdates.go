package services

import "log"

func UpdateStatus(status []string) {
	log.Printf("Updating status of item %s to %s", status[1], status[0])
}
