package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	if feedURL == "" {
		return &RSSFeed{}, fmt.Errorf("")
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agnet", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(dat, &feed)
	if err != nil {
		return &RSSFeed{}, err
	}

	return &feed, nil
}
