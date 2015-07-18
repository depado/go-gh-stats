package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var debug bool

func simpleReadLine() (l string, err error) {
	rl := bufio.NewScanner(os.Stdin)
	rl.Scan()
	if err = rl.Err(); err != nil {
		return
	}
	l = rl.Text()
	return
}

func main() {
	var err error

	flag.BoolVar(&debug, "debug", false, "Activates the debug logs")
	flag.Parse()

	if debug {
		log.Println("Started in debug mode")
	}
	fmt.Print("Enter a valid Github Username : ")
	usr, err := simpleReadLine()
	if err != nil {
		log.Fatalln("Could not read properly :", err)
	}
	if debug {
		log.Println("Fetching user data")
	}
	fmt.Println("Please wait while data are being retrieved...")
	usra, err := fetchUserData(usr)
	if err != nil {
		log.Fatalln("Could not fetch user data : ", err)
	}
	if debug {
		log.Println("Fetching and analysing repos.")
	}
	an, err := analyseRepos(usr)
	if err != nil {
		log.Fatalln("Error when analysing :", err)
	}
	fmt.Println(usr+"'s account exists since", usra.CreatedAt.String())
	fmt.Println()
	fmt.Println("Number of repositories :", an.nrepo)
	fmt.Println("Most starred repository :", an.mostStarred, "with", an.highestStars, "stars.")
	fmt.Println("Most forked repository :", an.mostForked, "with", an.highestForks, "forks.")
	fmt.Println("Most watched repository :", an.mostWatched, "with", an.highestWatches, "watchers.")
	fmt.Println()
	fmt.Println("Average stars per repository :", an.avStarsPerRepo)
	fmt.Println("Average forks per repository :", an.avForksPerRepo)
	fmt.Println("Average watchers per repository :", an.avWatchesPerRepo)
	fmt.Println()
	fmt.Println("Total stars :", an.totalStars)
	fmt.Println("Total forks :", an.totalForks)
	fmt.Println("Total watchers :", an.totalWatches)
}
