// combineHtml.go
// program that creates the directory structure for a domain
//
// author: prr, azul softwarw
// date: 14 August 2022
// copyright 2022 prr, azul software
//
package main

import (
    "os"
    "fmt"
	"bytes"
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

func findLinks(inbuf []byte)(links []string) {

	stylEnd := bytes.Index(inbuf, []byte("</style>"))

	fmt.Printf("stylEnd: %d\n", stylEnd)

	stPos := stylEnd+8
	linkNum := -1
	for i:=0; i<10; i++ {
		comPos := bytes.Index(inbuf[stPos:], []byte("<!--")) +stPos
		if comPos == stPos -1 {
			break
		}
		include := bytes.Index(inbuf[comPos:], []byte("include"))
		if include == -1 { continue}
		comEndPos := bytes.Index(inbuf[comPos:], []byte("-->"))
		if comEndPos == -1 {continue}
		comEndPos = comPos + comEndPos
		linkNum++
//		fmt.Printf("link: %d stPos: %d %s %s\n", linkNum, comPos, string(inbuf[(comPos + include + 7):comEndPos]),string(inbuf[comPos: comEndPos]))
		links = append(links, string(inbuf[(comPos + include + 7):comEndPos]))

		stPos = comPos + 5
	}


	return links
}

func main() {

    dbg = false
    site := ""
    numArgs := len(os.Args)

    switch numArgs {
        case 1:
            fmt.Println("no arguments provided!")
            fmt.Println("usage is:")
            fmt.Printf("%s htmlFile [dbg]\n", os.Args[0])
            os.Exit(-1)
        case 2:

            site = os.Args[1]
            if site == "dbg" {errmsg("invalid html filename: dbg!", nil)}
//            if site == "base" {errmsg("invalid site: base!", nil)}
//            fmt.Printf("** creating folders for site %s ***\n",site)

        case 3:
            site = os.Args[1]
            if site == "dbg" {errmsg("invalid html filename: dbg!", nil)}
//            if site == "base" {errmsg("invalid site: base!", nil)}
            fmt.Printf("** creating folders for site: %s ***\n",site)
            if os.Args[2] != "dbg" {
                fmt.Printf("error invalid argument: %s\n", os.Args[2])
                os.Exit(-1)
            }
            fmt.Println("dbg enabled!")
            dbg = true

        default:
            fmt.Println("too many commandline args!")
            fmt.Println("usage is:")
            fmt.Printf("%s htmlFile [dbg]\n", os.Args[0])
            os.Exit(-1)
    }

    // check whether baseline file
	extPos:=-1
	for i:=len(site)-1; i>-1; i-- {
		if site[i] == '.' {
			extPos = i
			break
		}
	}

	if extPos < 0 {
		fmt.Printf("no file name extension found!\n")
		os.Exit(-1)
	}

	extStr := string(site[extPos:])
	if extStr != ".html" {
		fmt.Printf("invalid file name extension: %s!\n", extStr)
		os.Exit(-1)
	}

	sitePath := "./inp/" + site
	fmt.Printf("parsing file: %s\n", sitePath)
//     := "/home/peter/www"

    fileInfo, err := os.Stat(sitePath)
    if os.IsNotExist(err) {
        fmt.Printf("error: file %s does not exist!\n", sitePath)
        os.Exit(-1)
    }
    if err != nil {
        fmt.Printf("error %v\n", err)
        os.Exit(-1)
    }

	size := int(fileInfo.Size())

//	fmt.Println("size: ", size)

	inbuf := make([]byte, size)

	infil, err := os.Open(sitePath)
	if err != nil {
		fmt.Printf("error opening %s: %v\n", sitePath, err)
		os.Exit(-1)
	}
	defer infil.Close()

	_, err = infil.Read(inbuf)
	if err != nil {
		fmt.Printf("error readin infil: %v\n", err)
		os.Exit(-1)
	}

	links := findLinks(inbuf)

	fmt.Printf("links: %d\n",len(links))
	for i:=0; i< len(links); i++ {
		fmt.Printf("link(%d): %s\n", i, links[i])
	}

	fmt.Println("*** success combine Html ***")
}
