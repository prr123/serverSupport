// rdSite.go
// program which reads yaml file and uses info to create a web site
//
// author: prr
// date: 9/8/2022
//
// copyright 2022 prr, azul software
//
package main

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v3"
)

type siteYamlObj struct {
	Domain string `yaml:"domain"`
	Desc string `yaml:"description"`
	Keys []string `yaml:"keywords"`
	Authors []string `yaml:"author"`
	Date string `yaml:"date"`
	Icon string `yaml:"icon"`
}



func dispYamlObj(site *siteYamlObj) {
	fmt.Println("*** Yaml Obj ***")
	fmt.Printf("  Domain: %s\n", site.Domain)
	fmt.Printf("  Description: %s\n", site.Desc)
	fmt.Printf("  keywords (%d):\n", len(site.Keys))
	for i:=0; i< len(site.Keys); i++ {
		fmt.Printf("    key %d: %s\n", i, site.Keys[i])
	}
	if len(site.Authors) == 1 {
		fmt.Printf("  author: %s\n", site.Authors[0])
	} else {
		fmt.Printf("  authors(%d):\n", len(site.Authors))
		for i:=0; i< len(site.Authors); i++ {
			fmt.Printf("    author(%d): %s\n", i, site.Authors[i])
		}
	}
	fmt.Printf("  date: %s\n", site.Date)
	fmt.Printf("  icon path: %s\n", site.Icon)
}

func checkYamlObj(site *siteYamlObj) (err error) {

	noErrParse := true
	errStr :=""

	if site == nil {return fmt.Errorf("siteYamlObj is nil!")}

	domain := site.Domain
	if !(len(domain)>0) {return fmt.Errorf("domain is empty!")}

	// check whether domain exists
	domainPath := "/home/peter/www/" + domain
	if !(len(domain) > 0) {
		noErrParse = false
		errStr += "domain does not exist!\n"
	} else {
	// if domin exists, check domain path
		_, err := os.Stat(domainPath)
		if err != nil {
			noErrParse = false
			errStr += fmt.Sprintf("domain '%s' not found: %v\n", domain, err)
		}
	}

	// description
	if !(len(site.Desc) > 0) {
		noErrParse = false
		errStr += "description does not exist!\n"
	}

	// keywords
	if !(len(site.Keys) > 0) {
		noErrParse = false
		errStr += "keywords do not exist!\n"
	}

	// author
	if !(len(site.Authors) > 0) {
		noErrParse = false
		errStr += "author(s) do not exist!\n"
	}

	// date
	if !(len(site.Date) > 0) {
		noErrParse = false
		errStr += "date does not exist!\n"
	}

	// icon
	if !(len(site.Icon) > 0) {
		noErrParse = false
		errStr += "icon does not exist!\n"
	} else {
	// if icon name exists, check icon path
		_, err := os.Stat(site.Icon)
		if err != nil {
			noErrParse = false
			errStr += fmt.Sprintf("icon file not found: %v\n", err)
		}
	}

	if !noErrParse {
		return fmt.Errorf("%s\n", errStr)
	}
	return nil
}

func main() {
	var yamlFilNam string

	numArgs:= len(os.Args)

	switch(numArgs) {
		case 1:
			fmt.Println("no yaml file name provided!")
			fmt.Println("usage is:")
			fmt.Println("rdSite 'yaml file name'")
			os.Exit(-1)


		case 2:
			yamlFilNam = os.Args[1] + ".yaml"

		default:
			fmt.Println("error - too many cmd line arguments!")
			fmt.Println("usage is:")
			fmt.Println("rdSite 'yaml file name'")
			os.Exit(-1)
	}

	fmt.Printf("yaml file: %s\n", yamlFilNam)

	yamlFilNamPath := "./inp/" + yamlFilNam
//	_, err := os.Stat(yamlFilNamPath)

	yamlFil, err := os.Open(yamlFilNamPath)
	if err != nil {
		fmt.Printf("error - opening yaml file: %v\n", err)
		os.Exit(-1)
	}
	defer yamlFil.Close()

	yamlFilInfo, err := yamlFil.Stat()
	size := int(yamlFilInfo.Size())
	yamlBuf := make([]byte,size)

	_, err = yamlFil.Read(yamlBuf)
	if err != nil {
		fmt.Printf("error - reading yaml data: %v\n", yamlBuf)
		os.Exit(-1)
	}

//	fmt.Printf("yaml data:\n %s\n", string(yamlBuf))

	yamlObj := new(siteYamlObj)
    err = yaml.Unmarshal(yamlBuf, yamlObj)
    if err != nil {
        fmt.Printf("unmarshall error: %v\n", err)
		os.Exit(-1)
    }
	fmt.Printf("yaml\n %v\n", yamlObj)

	dispYamlObj(yamlObj)

	err = checkYamlObj(yamlObj)
	if err != nil {
		fmt.Printf("\nerror parsing YamlObj:\n%v\n", err)
		os.Exit(-1)
	}

	fmt.Println("*** rdSite  success ***")
}
