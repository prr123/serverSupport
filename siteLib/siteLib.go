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
)

func creSite(site string) (err error) {

    // check whether baseline folder exists
    baseDir := "/home/peter/www"

    fileInfo, err := os.Stat(baseDir)
    if os.IsNotExist(err) {
        return fmt.Errorf("error: folder %s does not exist!", baseDir)
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
