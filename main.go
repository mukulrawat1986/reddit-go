// A test script to run our package
package main

import (
	"fmt"
	"github.com/mukulrawat1986/reddit-go/redditgo"
	"log"
)

func main() {
	// Call the Get function to get our structs filled with reddit data
	items, err := redditgo.Decode("aww")
	if err != nil {
		log.Fatal(err)
	}

	// Print the content of our Item struct
	for _, item := range items {
		fmt.Println(item)
	}

    for _, item := range items {
        _ = redditgo.Download_images(item.URL, item.Title)
    }
}
