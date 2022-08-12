// creHeader.go
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

type headerYamlObj struct {
	Domain string `yaml:"domain"`
	HtmlDir string `yaml:"htmlDir"`
	// Service
	MenuItems []string `yaml:"menuItems"`
	MenuRefs []string `yaml:"menuRefs"`
}

func dispHeaderYamlObj(header *headerYamlObj) {

	fmt.Println("*** Yaml Footer Obj ***")
	fmt.Printf("  domain:    %s\n", header.Domain)
	fmt.Printf("  htmlDir:   %s\n", header.HtmlDir)
	fmt.Printf("menu:\nitems(%d)  hrefs(%d)\n", len(header.MenuItems), len(header.MenuRefs))
	countItems := len(header.MenuItems)
	if len(header.MenuRefs) < countItems {
		countItems = len(header.MenuRefs)
		fmt.Println("error mismatch between MenuItems and MenuRefs!")
	}
	for i:=0; i< countItems; i++ {
		fmt.Printf("item %d: %-15s %-15s\n",i, header.MenuItems[i], header.MenuRefs[i])
	}

	return
}

func readYamlFil(yamlFil *os.File) (yamlHdObj *headerYamlObj, err error) {

	yamlFilInfo, err := yamlFil.Stat()
    size := int(yamlFilInfo.Size())
    yamlBuf := make([]byte,size)

    _, err = yamlFil.Read(yamlBuf)
    if err != nil {
        return nil, fmt.Errorf("error reading yaml data: %v\n", err)
    }

//  fmt.Printf("yaml data:\n %s\n", string(yamlBuf))

    yamlHdObj = new(headerYamlObj)
    err = yaml.Unmarshal(yamlBuf, yamlHdObj)
    if err != nil {
        return nil, fmt.Errorf("unmarshall error: %v\n", err)
    }

//	fmt.Printf("yaml:\n%v\n", yamlftObj)
	return yamlHdObj, nil
}

func creHeaderHtml(outfil *os.File, header *headerYamlObj) (err error) {

	outstr := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="description" content="a new test site">
  <meta name="keywords" content="key1, key2, key3">
  <meta name="author" content="prr, azul software">
  <meta name="date" content="1\8\2022">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Azul Software</title>
<link rel="icon" type="image/png" sizes="32x32" href="/home/peter/www/azul/image/azul32.png">
<style>
* {
  margin: 0;
  padding: 0;
  font-family: sans-serif;
  list-style: none;
  text-decoration:none;
}`
	outstr += `
.link {
  color: white;
  background-color: transparent;
  text-decoration: none;
}
.link:visited {
  color: pink;
  background-color: transparent;
  text-decoration: none;
}
.link:hover {
  color: red;
  background-color: transparent;
  text-decoration: underline;
}
`
	outstr +=`
.header {
	width:100%;
	height:100px;
	background-color: cyan;
	display: inline-flex;
	flex-wrap: wrap;
}
.col3 {
	width: 49.9%;
	background-color: pink;
	text-align: center;
}
.col1 {
	width: 24.9%;
}
.headerh2 {
	padding: 1em 0 0 1em;
}
`
	outstr += `
</style>
</head>
<body>
`

	outstr += "<section class=\"header\">\n"

	outstr += "  <div class=\"col1\">\n"
	outstr += "    <h2 class=\"headerh2\">LOGO</h2>\n"
	outstr += "  </div>\n"

	outstr += "  <div class=\"col3\">\n"
	outstr += "    <h2 class=\"headerh2\">THE GREAT STARTUP</h2>\n"
	outstr += "  </div>\n"

	outstr += "  <div class=\"col1\">\n"
	outstr += "    <h2 class=\"headerh2\">Social Media</h2>\n"
	outstr += "  </div>\n"

	outstr += "</section>"

	outstr += `
</body>
</html>
`

	_, err = outfil.WriteString(outstr)

	if err != nil {
		return fmt.Errorf("error writing footer html file: %v", err)
	}
	return nil
}

func main () {

	yamlFil, err := os.Open("./inp/header.yaml")
	if err != nil {
		fmt.Printf("error opening header yaml file: %v\n", err)
		os.Exit(-1)
	}
	defer yamlFil.Close()

	outfil, err := os.Create("./output/header.html")
	if err != nil {
		fmt.Printf("error creating header file: %v\n", err)
		os.Exit(-1)
	}
	defer outfil.Close()

	yamlObj, err := readYamlFil(yamlFil)
	if err != nil {
		fmt.Printf("error reading footer file: %v\n", err)
		os.Exit(-1)
	}

	dispHeaderYamlObj(yamlObj)

	err = creHeaderHtml(outfil, yamlObj)
	if err != nil {
		fmt.Printf("error creating html header file: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("*** success creating header ***")
}
