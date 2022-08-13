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
    "encoding/base64"
	"gopkg.in/yaml.v3"
	util "util"
)

type siteYamlObj struct {
	Domain string `yaml:"domain"`
	Desc string `yaml:"description"`
	Keys []string `yaml:"keywords"`
	Authors []string `yaml:"author"`
	Date string `yaml:"date"`
	Icon string `yaml:"icon"`
	Title string `yaml:"title"`
	PrimDom string
	SecDom string
}



func dispYamlObj(site *siteYamlObj) {
	fmt.Println("*** Yaml Obj ***")
	fmt.Printf("  Title:  %s\n", site.Title)
	fmt.Printf("  Domain: %s\n    Primary: %s Secondary: %s\n", site.Domain, site.PrimDom, site.SecDom)
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

func checkDomain(site *siteYamlObj)(err error) {

	db := []byte(site.Domain)
	point:= 0
	ptpos := 0
	for i:=0; i< len(db); i++ {
		if !util.IsAlphaNumeric(db[i]) {
			if db[i] == '.' {
				ptpos = i
				point++
			} else {
				return fmt.Errorf("char (%d): %d not alphaNumeric or period!", i)
			}
		}
	}

	if point == 0 {return fmt.Errorf("no subdomain!")}
	if point > 1 {return fmt.Errorf("more than one domain!")}

	site.SecDom = string(db[:ptpos])
	site.PrimDom = string(db[(ptpos+1):])

	return nil
}

func checkYamlObj(site *siteYamlObj) (err error) {

	noErrParse := true
	errStr :=""

	if site == nil {return fmt.Errorf("siteYamlObj is nil!")}

	// title
	title := site.Title
	if !(len(title)>0) {return fmt.Errorf("title is empty!")}

	// domain
	domain := site.Domain
	if !(len(domain)>0) {return fmt.Errorf("domain is empty!")}

	err = checkDomain(site)
	if err != nil {return fmt.Errorf("%v", err)}

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

func creIconData(iconFilNam string)(data string, err error){

	var base64Encoding string

	extPos := -1
	for i:=len(iconFilNam)-1; i>=0; i-- {
		if iconFilNam[i] == '.' {
			extPos = i
			break
		}
	}
	extStr := string(iconFilNam[(extPos+1):])

    switch extStr {
    case "jpeg", "jpg":
        base64Encoding += "data:image/jpeg;base64,"
    case "png":
        base64Encoding += "data:image/png;base64,"
	default:
		return "", fmt.Errorf("invalid mime %s!", extStr)
    }

	iconFil, err := os.Open(iconFilNam)
	if err != nil { return "", err }

	iconInfo, err := iconFil.Stat()
	if err != nil {return "", err}
	nb := int(iconInfo.Size())

	buf := make([]byte, nb)
	_,err = iconFil.Read(buf)
	if err != nil { return "", err }

	data = base64Encoding + base64.StdEncoding.EncodeToString(buf)
	return data, nil
}


func creIndexFil(site *siteYamlObj) (err error) {

	outFilPath := "./output/index.html"

	_, err = os.Stat(outFilPath)
	if err == nil {
		return fmt.Errorf("error index.html found!")
	} else {
		if os.IsExist(err){return fmt.Errorf("error %v!", err)}
	}

	fmt.Println("creating index.html!")
	outfil, err := os.Create(outFilPath)
	if err != nil {return fmt.Errorf("could not create index.html: %v", err)}
	defer outfil.Close()

	outstr := "<!DOCTYPE html>\n"
	outstr += "<html lang=\"en\">\n<head>\n"
	outstr += "  <meta charset=\"UTF-8\">\n"
	outstr += "  <meta name=\"description\" content=\""+site.Desc+"\">\n"
	outstr += "  <meta name=\"keywords\" content=\""
	keyStr:= ""
	for i:=0; i< (len(site.Keys) -1); i++ {
		keyStr += site.Keys[i] + ", "
	}
	if len(site.Keys) > 1 {keyStr += site.Keys[(len(site.Keys)-1)]}
	outstr += keyStr + "\">\n"
	outstr += "  <meta name=\"author\" content=\""
	authorStr := ""
	for i:=0; i< len(site.Authors) -1; i++ {
		authorStr += site.Authors[i] + ", "
	}
	if len(site.Authors) > 1 {authorStr += site.Authors[(len(site.Authors)-1)]}
	outstr += authorStr + "\">\n"
	outstr += "  <meta name=\"date\" content=\"" + site.Date + "\">\n"
	outstr += "  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n"

	outfil.WriteString(outstr)

	//link + title
	outstr = "<title>" + site.Title + "</title>\n"
//	outstr += "<link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"" + site.Icon + "\">\n"
	imgdata, err := creIconData(site.Icon)
	if err !=nil {return fmt.Errorf("creIconData: %v", err)}
	outstr += "<link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"" + imgdata + "\">\n"
	outfil.WriteString(outstr)

	outstr = `<style>
* {
  margin: 0;
  padding: 0;
  font-family: sans-serif;
  list-style: none;
  text-decoration:none;
}
</style>
</head>
<body>
`
	outfil.WriteString(outstr)

	outstr = `</body>
</html>
`
	outfil.WriteString(outstr)

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
//	fmt.Printf("yamlbU\n %v\n", yamlObj)
    err = yaml.Unmarshal(yamlBuf, yamlObj)
    if err != nil {
        fmt.Printf("unmarshall error: %v\n", err)
		os.Exit(-1)
    }
//	fmt.Printf("yamlaU\n %v\n", yamlObj)

	err = checkDomain(yamlObj)
	if err != nil {
        fmt.Printf("checkDomain: %v\n", err)
		os.Exit(-1)
	}
	dispYamlObj(yamlObj)

//	err = checkYamlObj(yamlObj)
	if err != nil {
		fmt.Printf("\nerror parsing YamlObj:\n%v\n", err)
		os.Exit(-1)
	}


// create index file
	err = creIndexFil(yamlObj)
	if err != nil {
		fmt.Printf("error creating Index File: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("*** creIdxFile  success ***")
}
