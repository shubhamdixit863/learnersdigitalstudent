package services

import (
	"time"
	"urlshortener/utils"
)

type URLService struct {
	urlMap     map[string]string
	expireTime map[string]time.Time
}

func NewURLService() *URLService {
	return &URLService{
		urlMap:    make(map[string]string),
		expireTime: make(map[string]time.Time),
	}
}

func (s *URLService) ShortenURL(url string) string {
	var shortCode string

	for {
		shortCode = utils.GenerateShortCode()
		if _, exists := s.urlMap[shortCode]; !exists {
			break
		}
	}

	s.urlMap[shortCode] = url
	s.expireTime[shortCode] = time.Now().Add(time.Hour) 

	return shortCode
}

func (s *URLService) GetURL(shortCode string) (string, bool) {

	url, exists := s.urlMap[shortCode]

	if !exists {
		return "", false
	}

	if time.Now().After(s.expireTime[shortCode]) {
		delete(s.urlMap, shortCode)
		delete(s.expireTime, shortCode)
		return "", false
	}

	return url, true
}

func (s *URLService) DeleteExpiredURL() {

	now := time.Now()

	for shortcode, expTime := range s.expireTime {
		if now.After(expTime) {
			delete(s.urlMap, shortcode)
			delete(s.expireTime, shortcode)
		}
	}

}

func (s *URLService) GetURLMap() map[string]time.Time{
	return s.expireTime
}