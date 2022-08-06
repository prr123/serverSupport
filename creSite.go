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
	numArgs := len(os.Args)

	switch numArgs {
		case 1:
			fmt.Println("no arguments provided!")

		case 2:
			if os.Args[1] != "dbg" {
				fmt.Printf("error invalid argument: %s\n", os.Args[1])
				os.Exit(-1)
			}
			dbg = true
		default:
			fmt.Println("too many args!")
			fmt.Println("usage is:")
			fmt.Printf("./%s \n", os.Args[0])
			os.Exit(-1)
	}

	fmt.Println("*** success ***")
}
