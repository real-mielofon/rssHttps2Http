package main

import (
	"fmt"
	"log"
	"github.com/real-mielofon/rss"
)

func main() {
	feed, err := rss.Fetch("http://feeds.rucast.net/Umputun")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(feed)
	
	err = feed.Update()
	if err != nil {
		log.Fatal(err)
	}
}