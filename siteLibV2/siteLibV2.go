// siteLib.go
// library that creates the directory structure for a domain
// and places index file into html
//
// author: prr, azul software
// date: 27 August 2022
// copyright 2022 prr, azul software
//
// folders
// root: /home/peter/www
// base: /home/peter/www/[base]
// sub
//       html
//       image
//       js
//       css
//
// v2 10/10/2022
// V2 add folders
//		svg
//		json
//		doc
//

package siteLibV2

import (
	"os"
	"fmt"
	"bufio"
    "encoding/base64"
    "io/ioutil"
    "net/http"
)


type IdxOpt struct {
	Favicon string
	Site string
	IdxFilNam string
}

var SubDirs =[7]string{"html","image","js","css","json","doc","svg"}


func ToBase64(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}

func GetExt (filNam string) (ext string) {
// function returns extension of a file name

	flen := len(filNam)
	extPos := -1
	for i:= flen-1; i>-1; i-- {
		if filNam[i] == '.' {
			extPos = i
			break
		}
	}
	if extPos == -1 {return ""}

	return string(filNam[extPos+1:]) 
}

func CreSite(site string) (err error) {

    // check whether baseline folder exists
    baseDir := "/home/peter/www"

    fileInfo, err := os.Stat(baseDir)
    if os.IsNotExist(err) {
        return fmt.Errorf("error: base directory (%s) does not exist!", baseDir)
    }
    if err != nil {	return err}

    if !fileInfo.IsDir() {
        return fmt.Errorf("error: %s is not a folder!", baseDir)
	}

    // now we need to check whether folder already exists
    siteDir := baseDir + "/" + site
    fileInfo, err = os.Stat(baseDir)
    if os.IsExist(err) {
        return fmt.Errorf("error: folder %s does exist!", siteDir)
    }

    err = os.Mkdir(siteDir, os.ModePerm)
    if err != nil {return fmt.Errorf("error creating folder: %v", err)}

    //creating subfolders and base files
	for i:=0; i<len(SubDirs); i++ {
		err = os.Mkdir(siteDir + "/" + SubDirs[i], os.ModePerm)
    	if err != nil {return fmt.Errorf("error creating html subfolder: %v", err)}
	}

	return nil
}

func DelSite(site string) (err error) {
// function deletes folders of the domain 'site'

    baseDir := "/home/peter/www"
    _, err = os.Stat(baseDir)
    if os.IsNotExist(err) {
        return fmt.Errorf("error: base directory (%s) does not exist!", baseDir)
    }
    if err != nil {	return err}

    // now we need to check whether folder already exists
    siteDir := baseDir + "/" + site
    _, err = os.Stat(baseDir)
    if os.IsExist(err) {
        fmt.Errorf("error: folder %s does exist!\n", siteDir)
        os.Exit(-1)
    }

    // make sure
    fmt.Printf("Deleting folders and files for site \"%s\"! Are you sure? (Y/n): ", site)
    reader := bufio.NewReader(os.Stdin)
    // ReadString will block until the delimiter is entered
    inp, err := reader.ReadString('\n')
    if err != nil {
        return fmt.Errorf("An error occured while reading input. Please try again: %v", err)
    }
    fmt.Printf("inp: %q\n", inp[0])
    if inp[0] != 'Y' {return fmt.Errorf("no correct confirmation!")}

    err = os.RemoveAll(siteDir)
    if err != nil {return fmt.Errorf("error deleting folder: %v", err)}

	return nil
}

func CreIndexFile(opt IdxOpt) (err error) {
// function creates an html index file
// opt file specifies: site, index file name, and favicon file

    wwwRoot := "/home/peter/www/"

    // check whether domain exists
    domainPath := wwwRoot + opt.Site

    _, err = os.Stat(domainPath)
    if err != nil {
        return fmt.Errorf("error - no domain \"%s\": %v", opt.Site, err)
    }

	idxFilNam:=""
    // check index file
    if len(opt.IdxFilNam) > 0 {
        idxFilNam = domainPath + "/html/" + opt.IdxFilNam + ".html"
    } else {
        idxFilNam = domainPath + "/html/index.html"
    }
    _, err = os.Stat(idxFilNam)
    if err == nil {
        return fmt.Errorf("error - index file exists: %s", idxFilNam)
    }

    fmt.Printf("*** creating index file for domain \"%s\" ***\n", opt.Site)
    idxFil, err := os.Create(idxFilNam)
    defer idxFil.Close()
    if err != nil {
        return fmt.Errorf("error creating idxFilNam! %v", err)
    }

	// test whether favicon file exists
	faviconPath := domainPath + "/image/" + opt.Favicon

//	faviconPath := "/home/peter/www/azul/image/azul32.png"

	favInfo, err := os.Stat(faviconPath)
    if os.IsNotExist(err) {
        return fmt.Errorf("error - favicon file does not exist: %s", faviconPath)
	}
	if err != nil {
        return fmt.Errorf("error - favicon file: %v", err)
	}

	ext := GetExt(faviconPath)
	imgTyp := false
	b64Typ := false
	switch ext {
		case "png64", "jpg64":
			b64Typ = true
		case "png", "jpg", "jpeg":
			imgTyp = true
		default:
			fmt.Printf("error -- favicon file extension %s not acceptable!\n", ext)
	}

	favStr :=""
	if b64Typ {
		favFil, err := os.Open(faviconPath)
		if err != nil { return fmt.Errorf("error opening faviconPath: %v", err)}
		defer favFil.Close()
		favBuf := make([]byte, favInfo.Size())
		_, err = favFil.Read(favBuf)
		if err != nil { return fmt.Errorf("error reading favicon: %v", err)}
		if ext == "png64" {
			favStr = `<link rel="icon" type="image/png" href="data:image/png;base64,`
		}
		if ext == "jpg64" {
			favStr = `<link rel="icon" type="image/jpeg" href="data:image/png;base64,`
		}
		favStr += string(favBuf[:]) + "\">\n"
	}
	if imgTyp {
		favStr = `<link rel="icon" type="image/` + ext
		favStr += `" sizes="32x32" href="`
		favStr += faviconPath + "\">\n"
	}
	if !(imgTyp || b64Typ) {
		fmt.Println("*** no favicon! ***")
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
  <title>Azul Software</title>`

	if b64Typ || imgTyp {outstr += favStr}

outstr += `<style>
* {
  margin: 0;
  padding: 0;
  font-family: sans-serif;
  list-style: none;
  text-decoration:none;
}

#main {
	border: 1px solid blue;
	min-height: 300px
}
</style>
</head>
<body>
 <script>azulLibV5.js</script>
 <script>azulindex.js</script>
</body>
</html>
`

    _, err = idxFil.WriteString(outstr)
    if err != nil {
        return fmt.Errorf("error writing: %v", err)
    }

	return nil
}

func EncodeB64 (filNam string) (err error) {

    var base64Encoding string

    inpFilNam := "./" + filNam

	// replace with read
    bytes, err := ioutil.ReadFile(inpFilNam)
    if err != nil {
        return fmt.Errorf("error no favicon file: %v", err)
    }

   // Determine the content type of the image file
    mimeType := http.DetectContentType(bytes)

    // Prepend the appropriate URI scheme header depending
    // on the MIME type
	ext :=""
    switch mimeType {
    case "image/jpeg":
        base64Encoding += "data:image/jpeg;base64,"
		ext = "jpg64"
    case "image/png":
        base64Encoding += "data:image/png;base64,"
		ext = "png64"
	default:
		return fmt.Errorf("not a valid mimetype: %s!", mimeType)
    }

    // Append the base64 encoded output
    base64Encoding += ToBase64(bytes)

    // Print the full base64 representation of the image
    extPs:=-1
    for i:=len(filNam) -1; i>-1;i-- {
        if filNam[i] == '.' {
            extPs = i
            break
        }
    }
    if extPs == -1 {extPs = len(filNam)}
    outFilNam := "./" + string(filNam[:extPs]) + ext
    outfil, err := os.Create(outFilNam)
    defer outfil.Close()
    if err != nil {
        return fmt.Errorf("error creating output file %s: %v\n", outFilNam, err)
    }

    outfil.WriteString(base64Encoding)

	return nil
}
