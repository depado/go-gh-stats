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

func debugln(t string) {
	if debug {
		log.Println(t)
	}
}

func perror(t string, err error) {
	if err != nil {
		log.Fatalln(t, err)
	}
}

func main() {
	var err error
	var usra userAPI
	var an analysis

	flag.BoolVar(&debug, "debug", false, "Activates the debug logs")
	flag.Parse()

	debugln("Started in debug mode")

	fmt.Print("Enter a valid Github Username : ")
	usr, err := simpleReadLine()
	perror("Could not read properly :", err)

	debugln("Fetching user data")
	fmt.Println("Please wait while data are being retrieved...")
	usra, err = fetchUserData(usr)
	perror("Could not fetch user data :", err)

	debugln("Fetching and analysing repos.")
	err = an.analyseRepos(usr)
	perror("Error when analysing :", err)

	debugln("Displaying analysis")
	usra.display()
	fmt.Println()
	an.display()
}
