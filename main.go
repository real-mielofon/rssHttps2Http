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
	
	//fmt.Println(feed)
	for _, item := range feed.Items {
		for _, en := range item.Enclosures {
			fmt.Println(en.URL)			
		}
	}
	
	err = feed.Update()
	if err != nil {
		log.Fatal(err)
	}
}