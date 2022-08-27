// creSite.go
// program that creates the directory structure for a domain
//
// author: prr, azul software
// date: 6 August 2022
// copyright 2022 prr, azul software
//
// folders
// root: /home/peter/www
// base: /home/peter/www/[base]
// sub
//       html
//       image
//       js
//       css

package main

import (
	"os"
	"fmt"
	site "mkSite/siteLib"
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

func helpfun() {

	fmt.Println("*** help ***")
	fmt.Println("usage is:")
	fmt.Printf("%s [site]/[help] [dbg]\n\n", os.Args[0])
	fmt.Printf("program creates subfolders for a web site\n")
	fmt.Printf("current default root folder is: '/home/peter/www'\n")
	fmt.Printf("base folder is:                 '/home/peter/www/[site]'\n")
	fmt.Printf("sub folders are:\n  html\n  image\n  css\n  js\n\n")
	os.Exit(0)
}

func main() {

	dbg = false
	help := false
	site := ""
	numArgs := len(os.Args)

	switch numArgs {
		case 1:
			fmt.Println("no arguments provided!")
			fmt.Println("usage is:")
			fmt.Printf("%s [site]/[help] [dbg]\n", os.Args[0])
			os.Exit(-1)
		case 2:

			site = os.Args[1]
			if site == "help" {help = true; break}
			if site == "dbg" {errmsg("invalid site: dbg!", nil)}
			if site == "base" {errmsg("invalid site: base!", nil)}
			fmt.Printf("*** creating folders for site %s ***\n",site)

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
			fmt.Printf("%s [site]/[help] [dbg]\n", os.Args[0])
			os.Exit(-1)
	}

	if help {
		helpfun()
		os.Exit(-1)
	}

	err = site.CreSite(site)
	if err != nil {
		fmt.Printf("error CreSite: %v!", err)
		os.Exit(-1)
	}

	fmt.Println("*** success ***")
}
