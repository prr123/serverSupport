// creSite.go
// program that creates the directory structure for a domain
//
// author: prr, azul softwarw
// date: 6 August 2022
// copyright 2022 prr, azul software
//
package main

import (
	"os"
	"fmt"
	"bufio"
)

var dbg bool

func errmsg(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	} else {
		fmt.Printf("%s\n", msg)
	}
	os.Exit(-1)
}

func main() {

	dbg = false
	site := ""
	numArgs := len(os.Args)

	switch numArgs {
		case 1:
			fmt.Println("no arguments provided!")
			fmt.Println("usage is:")
			fmt.Printf("./%s site [dbg]\n", os.Args[0])
			os.Exit(-1)
		case 2:
			site = os.Args[1]
			if site == "dbg" {errmsg("invalid site: dbg!", nil)}
			if site == "base" {errmsg("cannot delete site \"base\"!", nil)}
			fmt.Printf("*** deleting all files and folders for site %s ***\n",site)

		case 3:
			site = os.Args[1]
			if site == "dbg" {errmsg("invalid site: dbg!", nil)}
			if site == "base" {errmsg("cannot delete site \"base\"!", nil)}
			fmt.Printf("*** deleting all files and folders for site: %s ***\n",site)
			if os.Args[2] != "dbg" {
				fmt.Printf("error invalid argument: %s\n", os.Args[2])
				os.Exit(-1)
			}
			fmt.Println("dbg enabled!")
			dbg = true
		default:
			fmt.Println("too many args!")
			fmt.Println("usage is:")
			fmt.Printf("./%s site [dbg]\n", os.Args[0])
			os.Exit(-1)
	}

	// check whether baseline folder exists
	baseDir := "/home/peter/www"

	fileInfo, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		fmt.Printf("error: folder %s does not exist!\n", baseDir)
		os.Exit(-1)
	}
	if err != nil {
		fmt.Printf("error %v\n", err)
		os.Exit(-1)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("error: %s is not a folder!\n", baseDir)
		os.Exit(-1)
	}


	// now we need to check whether folder already exists
	siteDir := baseDir + "/" + site
	fileInfo, err = os.Stat(baseDir)
	if os.IsExist(err) {
		fmt.Printf("error: folder %s does exist!\n", siteDir)
		os.Exit(-1)
	}

	// make sure
	fmt.Printf("Deleting folders and files for site \"%s\"! Are you sure? (Y/n): ", site)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Printf("inp: %q\n", inp[0])
	if inp[0] != 'Y' {errmsg("no correct confirmation!", nil)}

	err = os.RemoveAll(siteDir)
	if err != nil {errmsg("error deleting folder", err)}

	fmt.Printf("*** success removed all folders from %s ***\n", siteDir)
}
