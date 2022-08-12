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
	Team string `yaml:"team"`
	Company string `yaml:"company"`
	Contact string `yaml:"contact"`
	// legal
	// privacy policy
	// data protection
	DataProtect string `yaml:"dataProtect"`
	// cookies
	Cookies string `yaml:"cookies"`
	// social media
	Fbook string `yaml:"facebook"`
	Twitter string `yaml:"twitter"`
	Linkedin string	`yaml:"linkedin"`
	Whatsapp string `yaml:"whatsapp"`
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

func creFooterHtml(outfil *os.File, footer *footerYamlObj) (err error) {

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
.col3 {
 width: 33.3%;
 background-color: black;
 color: white;
 padding: 10px 0 10px 0;
}
.footerul {
	padding-left: 1em;
}
.footerli {
  padding: 0.2em 0 0.2em 0;
}
.footerh2 {
    padding: 0.2em 0.2em 0.2em 0.5em;
}`
	outstr += `
</style>
</head>
<body>
`
	outstr += "<section class=\"footer\" style=\"width: 100%; display: inline-flex; flex-wrap: wrap;\">\n"

	outstr += "  <div class=\"col3\">\n"
	outstr += "    <h2 class=\"footerh2\">About Us </h2>\n"
	outstr += "    <ul class=\"footerul\">\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.AboutUs + "\">"+ "About Us</a></li>\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.Team + "\">"+ "Team</a></li>\n"
	outstr += "    </ul>\n"
	outstr += "  </div>\n"

	outstr += "  <div class=\"col3\">\n"
	outstr += "    <h2 class=\"footerh2\">Legal</h2>\n"
	outstr += "    <ul class=\"footerul\">\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.Company + "\">"+ "Organisation</a></li>\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.Cookies + "\">"+ "Cookies</li>\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.DataProtect + "\">"+ "Data Privacy</a></li>\n"
	outstr += "      <li class=\"footerli\"><a class=\"link\" href=\"" + footer.HtmlDir + "/" + footer.Contact + "\">"+ "Contact</a></li>\n"
	outstr += "    </ul>\n"
	outstr += "  </div>\n"

	outstr += "  <div class=\"col3\">\n"
	outstr += "    <h2 class=\"footerh2\">Social Media</h2>\n"
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

	err = creFooterHtml(outfil, yamlObj)
	if err != nil {
		fmt.Printf("error creating html footer file: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("*** success create footer ***")
}
