// creFooter.go
// program which reads yaml file and uses info to create a footer
//
// author: prr, azul software
// date: 11/8/2022
//
// copyright 2022 prr, azul software
//

package main

import (
	"fmt"
	"os"
//    "gopkg.in/yaml.v3"
//    util "util"
)

type footerYamlObj struct {
	// Service
	aboutUs string `yaml:"aboutUs"`
	contact string `yaml:"contact"`
	// legal
	// privacy policy
	// data protection
	dataProtect string `yaml:"dataProtect"`
	// cookies
	cookies string `yaml:"cookies"`
	// social media
	fbook string `yaml:"facebook"`
	twitter string `yaml:"twitter"`
	linkedin string	`yaml:"linkedin"`
	whatsapp string `yaml:"whatsapp"`

}

func main () {

	infil, err := os.Open("./inp/footer.yaml")
	if err != nil {
		fmt.Printf("error opening footer yaml file: %v\n", err)
		os.Exit(-1)
	}
	defer infil.Close()

	outfil, err := os.Create("./output/footer.html")
	if err != nil {
		fmt.Printf("error creating footer file: %v\n", err)
		os.Exit(-1)
	}
	defer outfil.Close()

	fmt.Println("*** success create footer ***")
}


