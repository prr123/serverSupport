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

package siteLib

import (
	"os"
	"fmt"
	"bufio"
)

func creSite(site string) (err error) {

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
    err = os.Mkdir(siteDir + "/html", os.ModePerm)
    if err != nil {return fmt.Errorf("error creating html subfolder: %v", err)}

    err = os.Mkdir(siteDir + "/image", os.ModePerm)
    if err != nil {return fmt.Errorf("error creating image subfolder: %v", err)}

    err = os.Mkdir(siteDir + "/js", os.ModePerm)
    if err != nil {return fmt.Errorf("error creating js subfolder: %v", err)}

    err = os.Mkdir(siteDir + "/css", os.ModePerm)
    if err != nil {return fmt.Errorf("error creating css subfolder: %v", err)}

	return nil
}

func delSite(site string) (err error) {

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
        fmt.Printf("error: folder %s does exist!\n", siteDir)
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

func creIndexFile(site string) (err error) {


	return nil
}
