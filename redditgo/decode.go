package redditgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Get(subreddit string, pageno string) ([]Item, error) {
	url := fmt.Sprintf("http://www.reddit.com/r/%s.json#page=%s", subreddit, pageno)
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}

	resp := new(Response)
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(resp.Data.Children))
	for i, child := range resp.Data.Children {
		items[i] = child.Data
	}
	//    after := resp.Data.After
	//    fmt.Printf("%s\n", after)
	return items, nil
}
