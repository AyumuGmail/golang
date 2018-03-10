package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch04/text/github"
)

func main() {
	result, err := github.SearchIsues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var resultsArrayClacified [3][]*github.Issue

	for _, item := range result.Items {
		//fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)

		oneMonthBefore := time.Now().AddDate(0, -2, 0)
		oneYearBefore := time.Now().AddDate(-1, 0, 0)

		if oneMonthBefore.Before(item.CreatedAt) {
			//append(classifiedResults[0], item)
			resultsArrayClacified[0] = append(resultsArrayClacified[0], item)
			//fmt.Println("*** 2ヶ月未満の投稿")
			//fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		if oneYearBefore.Before(item.CreatedAt) {
			resultsArrayClacified[1] = append(resultsArrayClacified[1], item)
			//append(classifiedResults[1], item)
			//fmt.Println("*** １年未満の投稿")
			//fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		if oneYearBefore.After(item.CreatedAt) {
			resultsArrayClacified[2] = append(resultsArrayClacified[2], item)
			//append(classifiedResults[2], item)
			//fmt.Println("*** １年以上の投稿")
			//fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}

	for i, classifiedResult := range resultsArrayClacified {
		if i == 0 {
			fmt.Println("**** 2ヶ月未満")
		} else if i == 1 {
			fmt.Println("**** 1年未満")
		} else if i == 2 {
			fmt.Println("**** 1年以上")
		}
		for _, item := range classifiedResult {
			fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
