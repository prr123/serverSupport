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

}

func dispYamlObj(site *siteYamlObj) {
	fmt.Println("*** Yaml Obj ***")
	fmt.Printf("  Domain: %s\n", site.Domain)
	fmt.Printf("  Description: %s\n", site.Desc)
	fmt.Printf("  keywords (%d):\n", len(site.Keys))
	for i:=0; i< len(site.Keys); i++ {
		fmt.Printf("key %d: %s\n", i, site.Keys[i])
	}

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
	fmt.Println("*** rdSite  success ***")
}
