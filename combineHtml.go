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
	"strings"
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

	if stylEnd == -1 {return links}
//	fmt.Printf("stylEnd: %d\n", stylEnd)

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
		linkStr := string(inbuf[(comPos + include + 7):comEndPos])
		linkStr = strings.TrimSpace(linkStr)
		links = append(links, linkStr)

		stPos = comPos + 5
	}


	return links
}

func insertLink(inbuf []byte, linkFilNam string)(outstr string, err error) {

	stylEnd := bytes.Index(inbuf, []byte("</style>"))
	if stylEnd == -1 {
		headPos := bytes.Index(inbuf, []byte("</head>"))
		if headPos == -1 {return "", fmt.Errorf("no head tag!")}
		outstr = string(inbuf[:headPos])
		outstr += "<style>\n"
	} else {
		outstr = string(inbuf[:stylEnd])
	}

	linkFilPath := "./output/" + linkFilNam
	filinfo, err1 := os.Stat(linkFilPath)
	if err1 != nil {return "", err1}

	size := int(filinfo.Size())

	linkfil, err1 := os.Open(linkFilPath)
	if err1 != nil {return "", err1}
	defer linkfil.Close()

	linkbuf := make([]byte, size)

	_, err1 = linkfil.Read(linkbuf)
	if err1 != nil {return "", err1}

	// discovering style
	stylPos := bytes.Index(linkbuf[:], []byte("<style>"))
	if stylPos == -1 {
		fmt.Println("link no style!")
	}

	linkSt:= -1
	stylEndPos := -1
	if stylPos > -1 {
		stylEndPos = bytes.Index(linkbuf[stylPos:], []byte("</style>"))
		if stylEndPos == -1 {return "", fmt.Errorf("link no /style tag found!")}
		linkSt = stylPos + 7

		istate :=0
		genStylEnd := 0
		for i:=stylPos; i< len(linkbuf); i++ {
			fmt.Printf("pos: %d char: %q istate: %d\n", i, linkbuf[i], istate)
			found := false
			switch istate {
				case 0:
					if linkbuf[i] == '*' {istate = 1}
				case 1:
					if linkbuf[i] == '{' {istate = 2; break;}
					if linkbuf[i] != ' ' { return "", fmt.Errorf("error parsing style -- found not whitespace! %q", linkbuf[i])}
				case 2:
					if linkbuf[i] == '}' {
						genStylEnd = i
						istate = 3
						found = true
					}
				default:
					found = true
			}
			if found {break}
		}
		if genStylEnd > 0 { linkSt = genStylEnd+1}
	}

	linkEnd := stylPos + stylEndPos

	stylStr := string(linkbuf[linkSt:linkEnd])
	fmt.Printf("stylstr:\n%s\n", stylStr)

	outstr += stylStr
	if stylEnd == -1 {
		outstr += "</style\n"
	}

	// find insert pos for html
	comPos := bytes.Index(inbuf[stylEnd:], []byte(linkFilNam))
	if comPos == -1 { return "", fmt.Errorf("link file name not found!")}
	comEndPos := bytes.Index(inbuf[stylEnd:], []byte("-->"))
	if comEndPos == -1 {return "", fmt.Errorf("link file end of comment not found!")}
	insPos := comEndPos + stylEnd + 4

	outstr += string(inbuf[stylEnd:insPos])

	// find html code in link file
	linkHtmlStPos := bytes.Index(linkbuf[linkEnd:], []byte("<body>"))
	if linkHtmlStPos == -1 { return "", fmt.Errorf("link file body tag not found!")}
	linkHtmlStPos += linkEnd
	linkHtmlEndPos := bytes.Index(linkbuf[linkHtmlStPos:], []byte("</body>"))
	if linkHtmlEndPos == -1 { return "", fmt.Errorf("link file /body tag not found!")}
	linkHtmlEndPos += linkHtmlStPos

	outstr += string(linkbuf[linkHtmlStPos:linkHtmlEndPos])

	outstr += string(inbuf[insPos:])
	return outstr, nil

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


	outFilNam := "./output/nindex.html"
	outfil, err := os.Create(outFilNam)
	if err != nil {
		fmt.Printf("error creating %s: %v\n", outFilNam, err)
		os.Exit(-1)
	}
	defer outfil.Close()

	links := findLinks(inbuf)

	outbuf := ""
	fmt.Printf("links: %d\n",len(links))
	for i:=0; i< len(links); i++ {
		fmt.Printf("\ninserting link(%d): %s\n", i, links[i])
		outbuf, err = insertLink(inbuf, links[i])
		if err != nil {
			fmt.Printf("error inserting link %s: %v\n", links[i], err)
			os.Exit(-1)
		}
		inbuf = make([]byte, len(outbuf))
		copy(inbuf, []byte(outbuf))
		fmt.Printf("merged link:\n%s\n\n", string(inbuf))
	}

	outfil.WriteString(outbuf)

	fmt.Println("*** success combining Html ***")
}
