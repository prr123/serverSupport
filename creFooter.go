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
    "gopkg.in/yaml.v3"
//    util "util"
)

type footerYamlObj struct {
	Domain string `yaml:"domain"`
	HtmlDir string `yaml:"htmlDir"`
	// Service
	AboutUs string `yaml:"aboutUs"`
	Contact string `yaml:"contact"`
	// legal
	// privacy policy
	// data protection
	DataProtect string `yaml:"dataProtect"`
	// cookies
	Cookies string `yaml:"cookies"`
	// social media
//	fbook string `yaml:"facebook"`
//	twitter string `yaml:"twitter"`
//	linkedin string	`yaml:"linkedin"`
//	whatsapp string `yaml:"whatsapp"`
}

func dispFooterYamlObj(footer *footerYamlObj) {

	fmt.Println("*** Yaml Footer Obj ***")
	fmt.Printf("  domain:    %s\n", footer.Domain)
	fmt.Printf("  htmlDir:   %s\n", footer.HtmlDir)
	fmt.Printf("  about us:  %s\n", footer.AboutUs)
	fmt.Printf("  contact:   %s\n", footer.Contact)
	fmt.Printf("  data prot: %s\n", footer.DataProtect)
	fmt.Printf("  cookies:   %s\n", footer.Cookies)

	return
}

func readYamlFil(yamlFil *os.File) (yamlftObj *footerYamlObj, err error) {

	yamlFilInfo, err := yamlFil.Stat()
    size := int(yamlFilInfo.Size())
    yamlBuf := make([]byte,size)

    _, err = yamlFil.Read(yamlBuf)
    if err != nil {
        return nil, fmt.Errorf("error reading yaml data: %v\n", err)
    }

//  fmt.Printf("yaml data:\n %s\n", string(yamlBuf))

    yamlftObj = new(footerYamlObj)
    err = yaml.Unmarshal(yamlBuf, yamlftObj)
    if err != nil {
        return nil, fmt.Errorf("unmarshall error: %v\n", err)
    }

//	fmt.Printf("yaml:\n%v\n", yamlftObj)
	return yamlftObj, nil
}


func main () {

	yamlFil, err := os.Open("./inp/footer.yaml")
	if err != nil {
		fmt.Printf("error opening footer yaml file: %v\n", err)
		os.Exit(-1)
	}
	defer yamlFil.Close()

	outfil, err := os.Create("./output/footer.html")
	if err != nil {
		fmt.Printf("error creating footer file: %v\n", err)
		os.Exit(-1)
	}
	defer outfil.Close()

	yamlObj, err := readYamlFil(yamlFil)
	if err != nil {
		fmt.Printf("error reading footer file: %v\n", err)
		os.Exit(-1)
	}
	dispFooterYamlObj(yamlObj)

	fmt.Println("*** success create footer ***")
}
