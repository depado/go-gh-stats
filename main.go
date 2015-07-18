package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var debug bool

// A simple readline operation. Reads stdin until Enter is pressed and returns
// the catched string.
func simpleReadLine() (l string, err error) {
	rl := bufio.NewScanner(os.Stdin)
	rl.Scan()
	if err = rl.Err(); err != nil {
		return
	}
	l = rl.Text()
	return
}

// Prints a string using log.Println if and only if the debug flag was passed.
// t : The string to be printed.
func debugln(t string) {
	if debug {
		log.Println(t)
	}
}

// Logs err if err isn't nil. Also exits the program.
func perror(t string, err error) {
	if err != nil {
		log.Fatalln(t, err)
	}
}

func main() {
	var err error

	var ua userAPI
	var an analysis

	flag.BoolVar(&debug, "debug", false, "Activates the debug logs")
	flag.Parse()

	debugln("Started in debug mode")

	fmt.Print("Enter a valid Github Username : ")
	usr, err := simpleReadLine()
	perror("Could not read properly :", err)

	fmt.Println("Please wait while data are being retrieved...")
	debugln("Fetching user data")
	err = ua.fetch(usr)
	perror("Could not fetch user data :", err)

	debugln("Fetching and analysing repos")
	err = an.analyseRepos(usr)
	perror("Error when analysing :", err)

	debugln("Displaying analysis")
	ua.display()
	fmt.Println() // Extra line break
	an.display()
}
