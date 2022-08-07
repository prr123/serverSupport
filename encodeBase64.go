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
	var filNam string

	numArgs := len(os.Args)
	switch numArgs {
		case 1:
			fmt.Println("no input file provided!")
			fmt.Println("  usage: encode file")
			os.Exit(-1)
		case 2:
			filNam = os.Args[1]
		default:
			fmt.Println("error too many args!")
			fmt.Println("  usage: encode file")
			os.Exit(-1)
	}

	inpFilNam := "./" + filNam

	bytes, err := ioutil.ReadFile(inpFilNam)
	if err != nil {
		fmt.Printf("error %v\n", err)
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
