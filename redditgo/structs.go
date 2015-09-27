package redditgo

import (
	"fmt"
)

type Item struct {
	Title   string
	URL     string
	Is_self bool
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
		After  string
		Before string
	}
}

func (i Item) String() string {
	return fmt.Sprintf("%s\n%s\n%v", i.Title, i.URL, i.Is_self)
}
