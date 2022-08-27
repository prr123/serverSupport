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
			fmt.Printf("%s site [dbg]\n", os.Args[0])
			os.Exit(-1)
		case 2:

			site = os.Args[1]
			if site == "dbg" {errmsg("invalid site: dbg!", nil)}
			if site == "base" {errmsg("invalid site: base!", nil)}
			fmt.Printf("** creating folders for site %s ***\n",site)

		case 3:
			site = os.Args[1]
			if site == "dbg" {errmsg("invalid site: dbg!", nil)}
			if site == "base" {errmsg("invalid site: base!", nil)}
			fmt.Printf("** creating folders for site: %s ***\n",site)
			if os.Args[2] != "dbg" {
				fmt.Printf("error invalid argument: %s\n", os.Args[2])
				os.Exit(-1)
			}
			fmt.Println("dbg enabled!")
			dbg = true
		default:
			fmt.Println("too many args!")
			fmt.Println("usage is:")
			fmt.Printf("%s site [dbg]\n", os.Args[0])
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

	err = os.Mkdir(siteDir, os.ModePerm)
	if err != nil {errmsg("error creating folder", err)}

	//creating subfolders and base files
	err = os.Mkdir(siteDir + "/html", os.ModePerm)
	if err != nil {errmsg("error creating html subfolder", err)}

	err = os.Mkdir(siteDir + "/image", os.ModePerm)
	if err != nil {errmsg("error creating image subfolder", err)}

	err = os.Mkdir(siteDir + "/js", os.ModePerm)
	if err != nil {errmsg("error creating js subfolder", err)}

	err = os.Mkdir(siteDir + "/css", os.ModePerm)
	if err != nil {errmsg("error creating css subfolder", err)}

	fmt.Println("*** success ***")
}
