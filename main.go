// A test script to run our package
package main

import (
	_ "fmt"
	"github.com/mukulrawat1986/reddit-go/redditgo"
	"log"
	"time"
)

func main() {
	// Call the Get function to get our structs filled with reddit data
	items, err := redditgo.Decode("aww")

	if err != nil {
		log.Fatalf("Error while decoding %v", err)
	}

	// Create a channel of booleans. The capacity of the channel will decide,
	// how many goroutines will be working at any point
	sema := make(chan bool, 5)

	// Download images from the link
	for _, item := range items {
		if !item.Is_self {
			go redditgo.Download_images(item.URL, item.Title, sema)
		} else {
			continue
		}
	}

	// Activate go routines
	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(40 * time.Second)
}
