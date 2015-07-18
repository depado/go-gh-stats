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
	usra, err := fetchUserData(usr)
	if err != nil {
		log.Fatalln("Could not fetch user data : ", err)
	}
	fmt.Println(usr+"'s account exists since", string(usra.CreatedAt.String()))
}
