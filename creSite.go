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
			if site == "dbg" {
				fmt.Println("invalid site: dbg!")
				os.Exit(-1)
			}
			fmt.Printf("** creating folders for site %s ***\n",site)

		case 3:
			site = os.Args[1]
			if site == "dbg" {
				fmt.Println("invalid site: dbg!")
				os.Exit(-1)
			}
			fmt.Printf("** creating folders for site %s ***\n",site)
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

	fmt.Println("*** success ***")
}
