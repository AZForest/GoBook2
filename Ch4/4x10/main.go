package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Exercise 4.10: Modify issues to report the results in age categories, say less than a month old,
// less than a year old, and more than a year old.

func main() {
	result, err := SearchIssues((os.Args[1:]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	tracker := make(map[string][]*Issue)

	for _, item := range result.Items {
		current := time.Now().Unix()
		var oneYear int64 = 31536000
		yearMarker := current - oneYear
		var oneMonth int64 = 2629743
		monthMarker := current - oneMonth

		itemSeconds := item.CreatedAt.Unix()
		

		//fmt.Printf("c: %d\ti: %d\tc: %d\tm: %d\top1: %v\top2: %v", current, itemSeconds, current, monthMarker, current - itemSeconds, current - monthMarker)
		if current - itemSeconds <= current - monthMarker {
			tracker["<1M"] = append(tracker["<1M"], item)
		} else if current - itemSeconds <= current - yearMarker {
			tracker["<1Y"] = append(tracker["<1Y"], item)
		} else {
			tracker[">1Y"] = append(tracker[">1Y"], item)
		}
	}
	displayItems(tracker)

}

func displayItems(tracker map[string][]*Issue) {
	fmt.Printf("Issues Less than 1 Month Old: \n")
	for _, item := range tracker["<1M"] {
		fmt.Printf("#%-5d %9.9s %d \n",
			item.Number, item.User.Login, item.CreatedAt.Year())
	}
	fmt.Printf("\nIssues Less than 1 Year Old: \n")
	for _, item := range tracker["<1Y"] {
		fmt.Printf("#%-5d %9.9s %d \n",
			item.Number, item.User.Login, item.CreatedAt.Year())
	}
	fmt.Printf("\nIssues Greater than 1 Year Old: \n")
	for _, item := range tracker[">1Y"] {
		fmt.Printf("#%-5d %9.9s %d \n",
			item.Number, item.User.Login, item.CreatedAt.Year())
	}
}

 
