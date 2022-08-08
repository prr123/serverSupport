package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	var filNam, domain, idxFilNam string

	wwwRoot := "/home/peter/www/"

	numArgs := len(os.Args)
	switch numArgs {
		case 1:
			fmt.Println("no domain provided!")
			fmt.Println("  usage: domain file")
			os.Exit(-1)
		case 2:
			domain = os.Args[1]

		case 3:
			domain = os.Args[1]
			filNam = os.Args[2]

		default:
			fmt.Println("error too many args!")
			fmt.Println("  usage: domain file")
			os.Exit(-1)
	}

	// check whether domain exists
	domainPath := wwwRoot + domain
	_, err := os.Stat(domainPath)
	if err != nil {
		fmt.Printf("error - no domain \"%s\": %v\n", domain, err)
		os.Exit(-1)
	}

	// check index file
	if len(filNam) > 0 {
		idxFilNam = domainPath + "/html/" + filNam + ".html"
	} else {
		idxFilNam = domainPath + "/html/index.html"
	}
	_, err = os.Stat(idxFilNam)
	if err == nil {
		fmt.Printf("error - index file exists: %s\n", idxFilNam)
		os.Exit(-1)
	}

	fmt.Printf("*** creating index file for domain \"%s\" ***\n", domain)
	idxFil, err := os.Create(idxFilNam)
	defer idxFil.Close()
	if err != nil {
		fmt.Printf("error creating idxFilNam! %v\n", err)
		os.Exit(-1)
	}

	outstr := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="description" content="blog about software writing and using">
  <meta name="keywords" content="Go">
  <meta name="author" content="prr">
  <meta name="date" content="1\3\2021">
  <meta  name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Azul Software</title>
<link rel="icon" type="image/png" sizes="32x32" href="/home/peter/www/azul/image/azul32.png">
<style>
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
  <h1>Hello from Server</h1>
</body>
</html>
`
	_, err = idxFil.WriteString(outstr)
	if err != nil {
		fmt.Printf("error writing: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("*** success ***")
	os.Exit(0)



	inpFilNam := "./" + filNam

	bytes, err := ioutil.ReadFile(inpFilNam)
	if err != nil {
		fmt.Printf("error no favicon file: %v\n", err)
		os.Exit(-1)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	extPs:=-1
	for i:=len(filNam) -1; i>-1;i-- {
		if filNam[i] == '.' {
			extPs = i
			break
		}
	}
	if extPs == -1 {extPs = len(filNam)}
	outFilNam := "./" + string(filNam[:extPs]) + ".b64"
	outfil, err := os.Create(outFilNam)
	defer outfil.Close()
	if err != nil {
		fmt.Printf("error creating output file %s: %v\n", outFilNam, err)
		os.Exit(-1)
	}
	outfil.WriteString(base64Encoding)

//	fmt.Println(base64Encoding)
}
