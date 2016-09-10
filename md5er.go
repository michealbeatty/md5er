package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//RecursiveHash will perform the MD5 operation recursively beginning with the
// directory in which the application is executed.
func recursiveHash(dir string, fout string) {
	f, err := os.Create(fout)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	filepath.Walk(dir, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				log.Print(err)
			}
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%x *%s\n", md5.Sum(fileData), rel)
			value := md5.Sum(fileData)

			s := hex.EncodeToString(value[:]) + " *" + rel + "\n"
			f.WriteString(s)
			//fmt.Println(path, file.Name())
		}
		return nil
	})
	f.Sync()
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	base := filepath.Base(cwd)
	fname := flag.String("o", base+".md5", "name of output file (should have md5 extension")
	flag.Parse()
	recursiveHash(cwd, *fname)
}

// TODO: add flag to allow user to specify if it should be recursive or not.
// If not recursive, should be only the contents of the current directory
// TODO: write results to file user should be able to specify a file name or default
// to root directory name. THis will use ioutil.WriteFile()
