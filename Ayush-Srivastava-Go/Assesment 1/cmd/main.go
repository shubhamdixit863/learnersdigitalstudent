package main

import (
	"fmt"
	"time"
	"urlshortener/services"
)

func main() {
	URLService := services.NewURLService()

	fmt.Println("Test 1: Shortening a URL...")
	url1 := "https://example.com"
	shortCode1 := URLService.ShortenURL(url1)
	fmt.Println("Shortened URL Code:", shortCode1)



	fmt.Println("\nTest 2: Retrieving Original URL...")
	retrievedURL, exists := URLService.GetURL(shortCode1)
	if exists {
		fmt.Println("Success: Retrieved URL ->", retrievedURL)
	} else {
		fmt.Println("Error: URL not found!")
	}




	fmt.Println("\nTest 3: Simulating URL Expiration...")
	shortCode2 := URLService.ShortenURL("https://expired.com")

	ExpireTime := URLService.GetURLMap()

	ExpireTime[shortCode2] = time.Now().Add(-time.Minute) 

	_, exists = URLService.GetURL(shortCode2)
	if !exists {
		fmt.Println("Success: Expired URL is not retrievable")
	} else {
		fmt.Println("Error: Expired URL should have been removed!")
	}



	fmt.Println("\nTest 4: Ensuring Unique Short Codes...")
	shortCode3 := URLService.ShortenURL("https://site1.com")
	shortCode4 := URLService.ShortenURL("https://site2.com")

	if shortCode3 != shortCode4 {
		fmt.Println("Success: Unique short codes generated!")
	} else {
		fmt.Println("Error: Duplicate short codes detected!")
	}




	fmt.Println("\nTest 5: Cleaning Up Expired URLs...")
	URLService.DeleteExpiredURL()
	fmt.Println("Expired URLs have been removed.")




	fmt.Println("\nTest 6: Retrieving Non-Existent Short Code...")
	_, exists = URLService.GetURL("invalidCode")
	if !exists {
		fmt.Println("Success: Handled non-existent short code correctly!")
	} else {
		fmt.Println("Error: Non-existent short code should not exist!")
	}

	fmt.Println("\nAll tests completed successfully!")
}
