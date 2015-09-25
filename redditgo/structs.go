package redditgo

import (
	"fmt"
)

type Item struct {
	Title string
	URL   string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func (i Item) String() string {
	return fmt.Sprintf("%s\n%s", i.Title, i.URL)
}
